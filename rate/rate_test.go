package rate

import (
	"errors"
	"testing"
	"time"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestNewSimpleLimiter(t *testing.T) {
	l := NewSimpleLimiter(3, 60)
	for i := 0; i < 3; i++ {
		r := l.Allow()
		AssertTrue(t, !r.Limited)
	}
	r := l.Allow()
	AssertTrue(t, r.Limited)
}

func TestNewLimiterDefault(t *testing.T) {
	l := NewLimiter(nil, nil)
	// 默认禁用，应该都放行
	r := l.Check("/api/test", "127.0.0.1")
	AssertTrue(t, !r.Limited)
}

func TestNewLimiterNilGlobal(t *testing.T) {
	// 仅配 IP，Global 为 nil 时不应 panic
	cfg := NewConfig(nil, &RateLimitItem{Limit: 2, Window: 60})
	l := NewLimiter(cfg, nil)
	AssertTrue(t, l.Config().Global == nil)

	AssertTrue(t, !l.CheckIP("1.1.1.1").Limited)
	AssertTrue(t, !l.CheckIP("1.1.1.1").Limited)
	r := l.CheckIP("1.1.1.1")
	AssertTrue(t, r.Limited)
	AssertEqual(t, r.Level, "ip")
}

func TestNewConfigNilDisablesLevel(t *testing.T) {
	cfg := NewConfig(nil, nil, RateLimitRouteItem{Route: "/api/sync", Limit: 2, Window: 60})
	AssertTrue(t, cfg.Enabled)
	AssertTrue(t, cfg.Global == nil)
	AssertTrue(t, cfg.PerIP == nil)
	AssertEqual(t, len(cfg.PerRoute), 1)

	l := NewLimiter(cfg, nil)
	// 综合检查不应触发默认的全局/IP 限额
	AssertTrue(t, !l.Check("/api/sync", "1.1.1.1").Limited)
	AssertTrue(t, !l.Check("/api/sync", "1.1.1.1").Limited)
	r := l.Check("/api/sync", "1.1.1.1")
	AssertTrue(t, r.Limited)
	AssertEqual(t, r.Level, "route")
}

func TestCheckGlobal(t *testing.T) {
	cfg := NewConfig(
		&RateLimitItem{Limit: 5, Window: 60},
		nil,
	)
	l := NewLimiter(cfg, nil)

	for i := 0; i < 5; i++ {
		r := l.CheckGlobal()
		AssertTrue(t, !r.Limited)
	}
	r := l.CheckGlobal()
	AssertTrue(t, r.Limited)
	AssertEqual(t, r.Level, "global")
}

func TestCheckIP(t *testing.T) {
	cfg := NewConfig(
		nil,
		&RateLimitItem{Limit: 3, Window: 60},
	)
	l := NewLimiter(cfg, nil)

	for i := 0; i < 3; i++ {
		r := l.CheckIP("10.0.0.1")
		AssertTrue(t, !r.Limited)
	}
	r := l.CheckIP("10.0.0.1")
	AssertTrue(t, r.Limited)
	AssertEqual(t, r.Level, "ip")

	// 其他 IP 不受影响
	r = l.CheckIP("10.0.0.2")
	AssertTrue(t, !r.Limited)
}

func TestCheckRoute(t *testing.T) {
	cfg := NewConfig(
		nil,
		nil,
		RateLimitRouteItem{Route: "/api/sync", Limit: 2, Window: 60},
	)
	l := NewLimiter(cfg, nil)

	_ = l.CheckRoute("/api/sync")
	r := l.CheckRoute("/api/sync")
	AssertTrue(t, !r.Limited) // 第 2 次
	r = l.CheckRoute("/api/sync")
	AssertTrue(t, r.Limited) // 第 3 次
	AssertEqual(t, r.Level, "route")

	// 未配置的路由放行
	r = l.CheckRoute("/api/other")
	AssertTrue(t, !r.Limited)
}

func TestCheckCombined(t *testing.T) {
	cfg := NewConfig(
		&RateLimitItem{Limit: 100, Window: 60},
		&RateLimitItem{Limit: 100, Window: 60},
		RateLimitRouteItem{Route: "/api/sync", Limit: 5, Window: 60},
	)
	l := NewLimiter(cfg, nil)

	// 连续 5 次：route 已达上限
	for i := 0; i < 5; i++ {
		r := l.Check("/api/sync", "1.1.1.1")
		AssertTrue(t, !r.Limited)
	}
	r := l.Check("/api/sync", "1.1.1.1")
	AssertTrue(t, r.Limited)
	AssertEqual(t, r.Level, "route")

	// 换 IP，仍然会 route 限制（因为按路由计数）
	r = l.Check("/api/sync", "2.2.2.2")
	AssertTrue(t, r.Limited)
	AssertEqual(t, r.Level, "route")
}

func TestConfigDisabled(t *testing.T) {
	cfg := DefaultConfig()
	l := NewLimiter(cfg, nil)

	for i := 0; i < 10000; i++ {
		r := l.Check("/api/sync", "1.1.1.1")
		AssertTrue(t, !r.Limited)
	}
}

func TestAllowN(t *testing.T) {
	l := NewSimpleLimiter(100, 60)

	for i := 0; i < 5; i++ {
		r := l.AllowN("custom-key", 5, 60)
		AssertTrue(t, !r.Limited)
	}
	r := l.AllowN("custom-key", 5, 60)
	AssertTrue(t, r.Limited)
}

func TestGetRouteRule(t *testing.T) {
	cfg := NewConfig(
		nil,
		nil,
		RateLimitRouteItem{Route: "/api/creator/sync", Limit: 5, Window: 60},
	)
	r := cfg.GetRouteRule("/api/creator/sync")
	AssertTrue(t, r != nil)
	AssertEqual(t, r.Limit, int64(5))

	r = cfg.GetRouteRule("/api/not-exist")
	AssertTrue(t, r == nil)

	// DefaultConfig 不含业务路由
	AssertTrue(t, DefaultConfig().GetRouteRule("/api/creator/sync") == nil)
}

func TestConfigAccess(t *testing.T) {
	l := NewSimpleLimiter(3, 60)
	AssertTrue(t, l.Config() != nil)
	AssertTrue(t, l.Store() != nil)
}

// errStore 始终返回错误，用于测试失败策略
type errStore struct{}

func (errStore) Incr(string, int64, int64) (CountResult, error) {
	return CountResult{}, errors.New("store unavailable")
}

func (errStore) Peek(string, int64, int64) (CountResult, error) {
	return CountResult{}, errors.New("store unavailable")
}

func (errStore) Cleanup() error { return nil }

func TestFailClosedOnStoreError(t *testing.T) {
	cfg := NewConfig(&RateLimitItem{Limit: 10, Window: 60}, nil)
	l := NewLimiter(cfg, errStore{}) // 默认 FailClosed
	r := l.Allow()
	AssertTrue(t, r.Limited)
	AssertEqual(t, r.Level, "error")

	r = l.AllowN("k", 10, 60)
	AssertTrue(t, r.Limited)
	AssertEqual(t, r.Level, "error")
}

func TestFailOpenOnStoreError(t *testing.T) {
	cfg := NewConfig(&RateLimitItem{Limit: 10, Window: 60}, nil)
	l := NewLimiter(cfg, errStore{}, WithFailOpen())
	r := l.Allow()
	AssertTrue(t, !r.Limited)
	AssertEqual(t, r.Level, "")
}

func TestCheckConsumesUpperLevels(t *testing.T) {
	// 路由限流触发后，全局配额仍已被消耗
	cfg := NewConfig(
		&RateLimitItem{Limit: 10, Window: 60},
		nil,
		RateLimitRouteItem{Route: "/api/sync", Limit: 1, Window: 60},
	)
	l := NewLimiter(cfg, nil)

	AssertTrue(t, !l.Check("/api/sync", "1.1.1.1").Limited)
	r := l.Check("/api/sync", "1.1.1.1")
	AssertTrue(t, r.Limited)
	AssertEqual(t, r.Level, "route")

	// 全局已被 Incr 两次（含被路由拒绝的那次）
	g := l.CheckGlobal()
	AssertTrue(t, !g.Limited)
	AssertEqual(t, g.Current, int64(3)) // 2 from Check + 1 from this CheckGlobal
}

func TestStartStopCleanup(t *testing.T) {
	l := NewSimpleLimiter(100, 60)
	l.StartCleanup(10 * time.Millisecond)
	l.StartCleanup(10 * time.Millisecond) // 重复启动应无效
	time.Sleep(25 * time.Millisecond)
	l.StopCleanup()
	l.StopCleanup() // 重复停止应安全
}

func TestWithCleanupInterval(t *testing.T) {
	l := NewSimpleLimiter(100, 60, WithCleanupInterval(10*time.Millisecond))
	time.Sleep(25 * time.Millisecond)
	l.StopCleanup()
}

func TestWithNonPositiveCleanupIntervalUsesDefault(t *testing.T) {
	for _, interval := range []time.Duration{0, -time.Second} {
		l := NewSimpleLimiter(100, 60, WithCleanupInterval(interval))
		AssertTrue(t, l.cleanupRunning)
		AssertEqual(t, l.cleanupInterval, time.Minute)
		l.StopCleanup()
	}
}

func TestPeekAndRemaining(t *testing.T) {
	l := NewSimpleLimiter(3, 60)

	r := l.PeekN("__global__", 3, 60)
	AssertTrue(t, !r.Limited)
	AssertEqual(t, r.Remaining, int64(3))
	AssertEqual(t, l.Remaining("__global__", 3, 60), int64(3))

	_ = l.Allow()
	_ = l.Allow()
	AssertEqual(t, l.Remaining("__global__", 3, 60), int64(1))

	r = l.PeekGlobal()
	AssertTrue(t, !r.Limited)
	AssertEqual(t, r.Current, int64(2))

	_ = l.Allow()
	r = l.PeekGlobal()
	AssertTrue(t, r.Limited)
	AssertEqual(t, r.Remaining, int64(0))

	// Peek 不消耗
	AssertEqual(t, l.PeekGlobal().Current, int64(3))
	AssertTrue(t, l.Allow().Limited)
}

func TestPeekCombined(t *testing.T) {
	cfg := NewConfig(
		nil,
		nil,
		RateLimitRouteItem{Route: "/api/sync", Limit: 1, Window: 60},
	)
	l := NewLimiter(cfg, nil)
	_ = l.CheckRoute("/api/sync")

	r := l.Peek("/api/sync", "1.1.1.1")
	AssertTrue(t, r.Limited)
	AssertEqual(t, r.Level, "route")

	// Peek 后计数不变，仍可再 Peek
	AssertEqual(t, l.PeekRoute("/api/sync").Current, int64(1))
}

func TestConfigCloneIsolation(t *testing.T) {
	cfg := NewConfig(&RateLimitItem{Limit: 5, Window: 60}, nil)
	l := NewLimiter(cfg, nil)

	cfg.Enabled = false
	cfg.Global.Limit = 1
	// 外部修改不影响 Limiter
	AssertTrue(t, l.Config().Enabled)
	AssertEqual(t, l.Config().Global.Limit, int64(5))

	snap := l.Config()
	snap.Enabled = false
	snap.Global.Limit = 1
	AssertTrue(t, l.Config().Enabled)
	AssertEqual(t, l.Config().Global.Limit, int64(5))
}

func TestGetRouteRuleReturnsCopy(t *testing.T) {
	cfg := NewConfig(nil, nil, RateLimitRouteItem{Route: "/api/sync", Limit: 5, Window: 60})
	r := cfg.GetRouteRule("/api/sync")
	r.Limit = 1
	AssertEqual(t, cfg.GetRouteRule("/api/sync").Limit, int64(5))
}
