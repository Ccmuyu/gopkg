package rate

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	. "github.com/Ccmuyu/gopkg/test"
)

// fakeRedis 用内存 ZSET 模拟本包 Lua 脚本的核心语义，便于单测。
type fakeRedis struct {
	mu   sync.Mutex
	data map[string]map[string]int64 // key -> member -> score
}

func newFakeRedis() *fakeRedis {
	return &fakeRedis{data: make(map[string]map[string]int64)}
}

func (f *fakeRedis) Eval(script string, keys []string, args ...any) (any, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if len(keys) != 1 {
		return nil, fmt.Errorf("want 1 key")
	}
	key := keys[0]
	now := toFakeInt(args[0])
	window := toFakeInt(args[1])
	limit := toFakeInt(args[2])
	cutoff := now - window

	z := f.data[key]
	if z == nil {
		z = make(map[string]int64)
		f.data[key] = z
	}
	for m, score := range z {
		if score <= cutoff {
			delete(z, m)
		}
	}

	isIncr := strings.Contains(script, "ZADD")
	var current int64
	if isIncr {
		member := fmt.Sprint(args[3])
		z[member] = now
		current = int64(len(z))
	} else {
		current = int64(len(z))
	}

	var limited int64
	if isIncr {
		if current > limit {
			limited = 1
		}
	} else if current >= limit {
		limited = 1
	}

	remaining := limit - current
	if remaining < 0 {
		remaining = 0
	}

	var resetAfter int64
	if limited == 1 && len(z) > 0 {
		oldest := int64(1<<63 - 1)
		for _, score := range z {
			if score < oldest {
				oldest = score
			}
		}
		resetAfter = oldest + window - now
		if resetAfter < 1 {
			resetAfter = 1
		}
	}

	return []int64{current, limited, resetAfter, remaining}, nil
}

func toFakeInt(v any) int64 {
	switch n := v.(type) {
	case int64:
		return n
	case int:
		return int64(n)
	default:
		panic(fmt.Sprintf("bad arg %T", v))
	}
}

func TestRedisStoreIncrAndPeek(t *testing.T) {
	fake := newFakeRedis()
	s := NewRedisStore(fake, "t:").(*redisStore)
	base := time.Unix(1_700_000_000, 0)
	s.nowFn = func() time.Time { return base }

	for i := 0; i < 3; i++ {
		r, err := s.Incr("k", 3, 60)
		AssertNoError(t, err)
		AssertTrue(t, !r.Limited)
	}
	r, err := s.Incr("k", 3, 60)
	AssertNoError(t, err)
	AssertTrue(t, r.Limited)
	AssertTrue(t, r.ResetAfter > 0)
	AssertEqual(t, r.Remaining, int64(0))

	r, err = s.Peek("k", 3, 60)
	AssertNoError(t, err)
	AssertTrue(t, r.Limited)
	AssertEqual(t, r.Current, int64(4))
}

func TestRedisStorePeekNoConsume(t *testing.T) {
	fake := newFakeRedis()
	s := NewRedisStore(fake, "")
	for i := 0; i < 2; i++ {
		_, err := s.Incr("u", 5, 60)
		AssertNoError(t, err)
	}
	r1, err := s.Peek("u", 5, 60)
	AssertNoError(t, err)
	r2, err := s.Peek("u", 5, 60)
	AssertNoError(t, err)
	AssertEqual(t, r1.Current, r2.Current)
	AssertEqual(t, r1.Current, int64(2))
	AssertEqual(t, r1.Remaining, int64(3))
}

func TestRedisStoreCleanupNoop(t *testing.T) {
	s := NewRedisStore(newFakeRedis(), "x:")
	AssertNoError(t, s.Cleanup())
}

func TestRedisStoreInvalidLimit(t *testing.T) {
	s := NewRedisStore(newFakeRedis(), "")
	r, err := s.Incr("k", 0, 60)
	AssertNoError(t, err)
	AssertTrue(t, r.Limited)
}

func TestParseRedisCountResult(t *testing.T) {
	r, err := parseRedisCountResult([]any{int64(2), int64(0), int64(0), int64(3)})
	AssertNoError(t, err)
	AssertEqual(t, r.Current, int64(2))
	AssertTrue(t, !r.Limited)
	AssertEqual(t, r.Remaining, int64(3))

	r, err = parseRedisCountResult([]int64{5, 1, 10, 0})
	AssertNoError(t, err)
	AssertTrue(t, r.Limited)
	AssertEqual(t, r.ResetAfter, int64(10))

	_, err = parseRedisCountResult("bad")
	AssertTrue(t, err != nil)
}

func TestNewRedisStoreNilPanics(t *testing.T) {
	defer func() {
		AssertTrue(t, recover() != nil)
	}()
	_ = NewRedisStore(nil, "")
}
