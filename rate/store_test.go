package rate

import (
	"testing"
	"time"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestSlidingWindowIncr(t *testing.T) {
	s := NewSlidingWindowStore(10, 5) // 10s 窗口, 5 槽 (每槽 2s)
	for i := 0; i < 3; i++ {
		r, err := s.Incr("k", 5, 10)
		AssertNoError(t, err)
		AssertTrue(t, !r.Limited)
	}
	r, err := s.Incr("k", 5, 10)
	AssertNoError(t, err)
	AssertTrue(t, !r.Limited)
}

func TestSlidingWindowLimit(t *testing.T) {
	s := NewSlidingWindowStore(10, 5)
	for i := 0; i < 3; i++ {
		_, _ = s.Incr("k", 3, 10)
	}
	r, err := s.Incr("k", 3, 10) // 第 4 次，超限
	AssertNoError(t, err)
	AssertTrue(t, r.Limited)
	AssertTrue(t, r.ResetAfter > 0)
}

func TestSlidingWindowExpire(t *testing.T) {
	// 用 mock 时间测试窗口滑动
	base := time.Unix(1000000, 0)
	offset := int64(0)
	nowFn := func() time.Time {
		return base.Add(time.Duration(offset) * time.Second)
	}
	s := NewSlidingWindowStore(10, 5).(*slidingWindowStore)
	s.nowFunc = nowFn

	// t=0~1 秒内发 3 次请求
	for i := 0; i < 3; i++ {
		r, err := s.Incr("k", 5, 10)
		AssertNoError(t, err)
		AssertTrue(t, !r.Limited)
	}

	// 时间推进 15 秒：原请求全部超出窗口
	offset = 15

	// 新的请求应该从 1 开始
	r, err := s.Incr("k", 5, 10)
	AssertNoError(t, err)
	AssertTrue(t, !r.Limited)
	AssertEqual(t, r.Current, int64(1))
}

func TestSlidingWindowCleanup(t *testing.T) {
	base := time.Unix(1000000, 0)
	offset := int64(0)
	nowFn := func() time.Time {
		return base.Add(time.Duration(offset) * time.Second)
	}
	// idleTTL = max(2*window, 120) = 120
	s := NewSlidingWindowStore(10, 5).(*slidingWindowStore)
	s.nowFunc = nowFn

	_, _ = s.Incr("k1", 10, 10)
	_, _ = s.Incr("k2", 10, 10)
	AssertEqual(t, s.lenKeys(), 2)

	// 未到空闲阈值，不应清理
	offset = 60
	AssertNoError(t, s.Cleanup())
	AssertEqual(t, s.lenKeys(), 2)

	// 超过 idleTTL=120，应清理
	offset = 200
	AssertNoError(t, s.Cleanup())
	AssertEqual(t, s.lenKeys(), 0)
}

func TestSlidingWindowCleanupUsesLastWindow(t *testing.T) {
	base := time.Unix(1000000, 0)
	offset := int64(0)
	nowFn := func() time.Time {
		return base.Add(time.Duration(offset) * time.Second)
	}
	// 构造 idleTTL=120，但 key 使用 window=300 → 空闲阈值应为 600
	s := NewSlidingWindowStore(10, 5).(*slidingWindowStore)
	s.nowFunc = nowFn

	_, _ = s.Incr("long", 100, 300)
	offset = 400 // < 600，不应清理
	AssertNoError(t, s.Cleanup())
	AssertEqual(t, s.lenKeys(), 1)

	offset = 700 // > 600，应清理
	AssertNoError(t, s.Cleanup())
	AssertEqual(t, s.lenKeys(), 0)
}

func TestSlidingWindowMultipleKeys(t *testing.T) {
	s := NewSlidingWindowStore(10, 5)

	for i := 0; i < 3; i++ {
		r, err := s.Incr("a", 3, 10)
		AssertNoError(t, err)
		AssertTrue(t, !r.Limited)
	}

	r, err := s.Incr("a", 3, 10)
	AssertNoError(t, err)
	AssertTrue(t, r.Limited)

	r, err = s.Incr("b", 3, 10)
	AssertNoError(t, err)
	AssertTrue(t, !r.Limited)
}

func TestSlidingWindowSlotRotation(t *testing.T) {
	base := time.Unix(1000000, 0)
	offset := int64(0)
	nowFn := func() time.Time {
		return base.Add(time.Duration(offset) * time.Second)
	}
	s := NewSlidingWindowStore(10, 5).(*slidingWindowStore)
	s.nowFunc = nowFn

	// 在不同的槽内请求
	_, _ = s.Incr("k", 100, 10)
	offset = 3
	_, _ = s.Incr("k", 100, 10)
	offset = 7
	_, _ = s.Incr("k", 100, 10)
	offset = 12
	r, _ := s.Incr("k", 100, 10)
	AssertTrue(t, !r.Limited)
	// 当前 t=12，窗口范围是 [2, 12]，最早在 t=0 的请求应已过期
	// 剩余计数应该来自 t=3,7,12：共 3 次
	AssertEqual(t, r.Current, int64(3))
}

func TestSlidingWindowNoPrematureDrop(t *testing.T) {
	// 即使跨越的槽数超过 slots，也不应丢弃仍在窗口内的计数
	base := time.Unix(1000000, 0)
	offset := int64(0)
	nowFn := func() time.Time {
		return base.Add(time.Duration(offset) * time.Second)
	}
	s := NewSlidingWindowStore(100, 5).(*slidingWindowStore) // slots=5
	s.nowFunc = nowFn

	// 每 1 秒一次，共 10 次（超过 slots=5）
	for i := 0; i < 10; i++ {
		offset = int64(i)
		r, err := s.Incr("k", 10, 100)
		AssertNoError(t, err)
		AssertTrue(t, !r.Limited)
		AssertEqual(t, r.Current, int64(i+1))
	}

	r, err := s.Incr("k", 10, 100)
	AssertNoError(t, err)
	AssertTrue(t, r.Limited)
	AssertEqual(t, r.Current, int64(10))
}

func TestSlidingWindowKeepsPartiallyOverlappingSlot(t *testing.T) {
	base := time.Unix(1_000_000, 0)
	offset := int64(4)
	s := NewSlidingWindowStore(10, 2).(*slidingWindowStore) // 每槽 5 秒
	s.nowFunc = func() time.Time {
		return base.Add(time.Duration(offset) * time.Second)
	}

	// 请求发生在槽 [0,5) 的末尾。
	_, _ = s.Incr("k", 10, 10)
	offset = 11 // 窗口起点为 1，请求仍在最近 10 秒内

	r, err := s.Peek("k", 10, 10)
	AssertNoError(t, err)
	AssertEqual(t, r.Current, int64(1))
}

func TestSlidingWindowRejectedRequestsDoNotConsumeQuota(t *testing.T) {
	base := time.Unix(1_000_000, 0)
	offset := int64(0)
	s := NewSlidingWindowStore(10, 5).(*slidingWindowStore)
	s.nowFunc = func() time.Time {
		return base.Add(time.Duration(offset) * time.Second)
	}

	for i := 0; i < 3; i++ {
		r, err := s.Incr("k", 3, 10)
		AssertNoError(t, err)
		AssertTrue(t, !r.Limited)
	}
	for i := 0; i < 100; i++ {
		r, err := s.Incr("k", 3, 10)
		AssertNoError(t, err)
		AssertTrue(t, r.Limited)
		AssertEqual(t, r.Current, int64(3))
	}

	// 最早一批所在的聚合槽过期后应立即恢复，不受此前 100 次拒绝影响。
	offset = 12
	r, err := s.Incr("k", 3, 10)
	AssertNoError(t, err)
	AssertTrue(t, !r.Limited)
	AssertEqual(t, r.Current, int64(1))
}

func TestSlidingWindowSameKeyDifferentWindowsAreIsolated(t *testing.T) {
	base := time.Unix(1_000_000, 0)
	s := NewSlidingWindowStore(60, 10).(*slidingWindowStore)
	s.nowFunc = func() time.Time { return base }

	for i := 0; i < 3; i++ {
		_, _ = s.Incr("k", 3, 60)
	}

	short, err := s.Incr("k", 10, 10)
	AssertNoError(t, err)
	AssertEqual(t, short.Current, int64(1))

	long, err := s.Peek("k", 3, 60)
	AssertNoError(t, err)
	AssertTrue(t, long.Limited)
	AssertEqual(t, long.Current, int64(3))
}

func TestSlidingWindowDifferentWindows(t *testing.T) {
	// 同一 Store 上使用不同 window 应各自按入参窗口计数
	base := time.Unix(1000000, 0)
	offset := int64(0)
	nowFn := func() time.Time {
		return base.Add(time.Duration(offset) * time.Second)
	}
	s := NewSlidingWindowStore(60, 10).(*slidingWindowStore)
	s.nowFunc = nowFn

	for i := 0; i < 3; i++ {
		r, err := s.Incr("short", 3, 10)
		AssertNoError(t, err)
		AssertTrue(t, !r.Limited)
	}
	r, err := s.Incr("short", 3, 10)
	AssertNoError(t, err)
	AssertTrue(t, r.Limited)

	// 另一 key、更大窗口不受影响
	r, err = s.Incr("long", 100, 60)
	AssertNoError(t, err)
	AssertTrue(t, !r.Limited)
	AssertEqual(t, r.Current, int64(1))
}

func TestSlidingWindowInvalidLimit(t *testing.T) {
	s := NewSlidingWindowStore(10, 5)
	r, err := s.Incr("k", 0, 10)
	AssertNoError(t, err)
	AssertTrue(t, r.Limited)
}

func TestSlidingWindowPeek(t *testing.T) {
	s := NewSlidingWindowStore(10, 5)

	r, err := s.Peek("k", 3, 10)
	AssertNoError(t, err)
	AssertTrue(t, !r.Limited)
	AssertEqual(t, r.Current, int64(0))
	AssertEqual(t, r.Remaining, int64(3))

	_, _ = s.Incr("k", 3, 10)
	_, _ = s.Incr("k", 3, 10)
	_, _ = s.Incr("k", 3, 10)

	r, err = s.Peek("k", 3, 10)
	AssertNoError(t, err)
	AssertTrue(t, r.Limited) // current==3，下一次 Incr 会被限
	AssertEqual(t, r.Current, int64(3))
	AssertEqual(t, r.Remaining, int64(0))

	// Peek 不消耗配额
	r, err = s.Peek("k", 3, 10)
	AssertNoError(t, err)
	AssertEqual(t, r.Current, int64(3))
}
