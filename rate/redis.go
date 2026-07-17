package rate

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"sync/atomic"
)

// RedisEvaler 执行 Redis EVAL。可用 go-redis 等适配，例如：
//
//	type adapter struct{ rdb *redis.Client }
//	func (a adapter) Eval(script string, keys []string, args ...any) (any, error) {
//	    return a.rdb.Eval(ctx, script, keys, args...).Result()
//	}
type RedisEvaler interface {
	Eval(script string, keys []string, args ...any) (any, error)
}

// redis 滑动窗口（ZSET）：按 score=unix 秒排序，窗口外成员用 ZREMRANGEBYSCORE 剔除。
const redisIncrScript = `
local key = KEYS[1]
local redisTime = redis.call('TIME')
local now = tonumber(redisTime[1])
local window = tonumber(ARGV[1])
local limit = tonumber(ARGV[2])
local member = ARGV[3]

redis.call('ZREMRANGEBYSCORE', key, '-inf', now - window)
local current = redis.call('ZCARD', key)

local limited = 0
if current >= limit then
  limited = 1
else
  redis.call('ZADD', key, now, member)
  current = current + 1
end
redis.call('EXPIRE', key, window)

local remaining = limit - current
if remaining < 0 then remaining = 0 end

local resetAfter = 0
if limited == 1 then
  local oldest = redis.call('ZRANGE', key, 0, 0, 'WITHSCORES')
  if #oldest >= 2 then
    resetAfter = tonumber(oldest[2]) + window - now
    if resetAfter < 1 then resetAfter = 1 end
  else
    resetAfter = 1
  end
end

return {current, limited, resetAfter, remaining}
`

const redisPeekScript = `
local key = KEYS[1]
local redisTime = redis.call('TIME')
local now = tonumber(redisTime[1])
local window = tonumber(ARGV[1])
local limit = tonumber(ARGV[2])

redis.call('ZREMRANGEBYSCORE', key, '-inf', now - window)
local current = redis.call('ZCARD', key)

local limited = 0
if current >= limit then limited = 1 end

local remaining = limit - current
if remaining < 0 then remaining = 0 end

local resetAfter = 0
if limited == 1 then
  local oldest = redis.call('ZRANGE', key, 0, 0, 'WITHSCORES')
  if #oldest >= 2 then
    resetAfter = tonumber(oldest[2]) + window - now
    if resetAfter < 1 then resetAfter = 1 end
  else
    resetAfter = 1
  end
end

return {current, limited, resetAfter, remaining}
`

type redisStore struct {
	eval       RedisEvaler
	prefix     string
	instanceID string
	seq        uint64
}

// NewRedisStore 创建基于 Redis 的滑动窗口 Store。
// evaler 不可为 nil；keyPrefix 为空时使用 "rate:"。
// Cleanup 为 no-op（依赖 key 的 EXPIRE）。
func NewRedisStore(evaler RedisEvaler, keyPrefix string) Store {
	if evaler == nil {
		panic("rate: RedisEvaler is nil")
	}
	if keyPrefix == "" {
		keyPrefix = "rate:"
	}
	return &redisStore{
		eval:       evaler,
		prefix:     keyPrefix,
		instanceID: newRedisInstanceID(),
	}
}

func newRedisInstanceID() string {
	var id [12]byte
	if _, err := rand.Read(id[:]); err != nil {
		panic(fmt.Sprintf("rate: generate redis instance id: %v", err))
	}
	return hex.EncodeToString(id[:])
}

func (s *redisStore) fullKey(key string, window int64) string {
	return fmt.Sprintf("%s%s:w:%d", s.prefix, key, window)
}

func (s *redisStore) Incr(key string, limit, window int64) (CountResult, error) {
	if limit <= 0 || window <= 0 {
		return CountResult{Limited: true}, nil
	}
	member := fmt.Sprintf("%s-%d", s.instanceID, atomic.AddUint64(&s.seq, 1))
	raw, err := s.eval.Eval(redisIncrScript, []string{s.fullKey(key, window)}, window, limit, member)
	if err != nil {
		return CountResult{}, err
	}
	return parseRedisCountResult(raw)
}

func (s *redisStore) Peek(key string, limit, window int64) (CountResult, error) {
	if limit <= 0 || window <= 0 {
		return CountResult{Limited: true}, nil
	}
	raw, err := s.eval.Eval(redisPeekScript, []string{s.fullKey(key, window)}, window, limit)
	if err != nil {
		return CountResult{}, err
	}
	return parseRedisCountResult(raw)
}

func (s *redisStore) Cleanup() error { return nil }

func parseRedisCountResult(raw any) (CountResult, error) {
	vals, err := toInt64Slice(raw)
	if err != nil {
		return CountResult{}, err
	}
	if len(vals) < 4 {
		return CountResult{}, fmt.Errorf("rate: unexpected redis result length %d", len(vals))
	}
	return CountResult{
		Current:    vals[0],
		Limited:    vals[1] != 0,
		ResetAfter: vals[2],
		Remaining:  vals[3],
	}, nil
}

func toInt64Slice(raw any) ([]int64, error) {
	switch v := raw.(type) {
	case []int64:
		return v, nil
	case []any:
		out := make([]int64, len(v))
		for i, x := range v {
			n, err := toInt64(x)
			if err != nil {
				return nil, err
			}
			out[i] = n
		}
		return out, nil
	default:
		return nil, fmt.Errorf("rate: unexpected redis result type %T", raw)
	}
}

func toInt64(v any) (int64, error) {
	switch n := v.(type) {
	case int64:
		return n, nil
	case int:
		return int64(n), nil
	case float64:
		return int64(n), nil
	default:
		return 0, fmt.Errorf("rate: cannot convert %T to int64", v)
	}
}
