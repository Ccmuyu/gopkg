package rate

import (
	"sync"
	"time"
)

// CheckResult 综合检查结果（业务方最常用）
type CheckResult struct {
	Limited    bool   // 是否被任一规则限流（或 FailClosed 下 Store 出错）
	ResetAfter int64  // 建议重试等待秒数
	Level      string // "global" / "ip" / "route" / "error" / ""
	Current    int64  // 当前窗口计数
	Remaining  int64  // 窗口内剩余可请求次数（综合 Check/Peek 时为触发级别的剩余）
}

// Option 限流器可选配置
type Option func(*Limiter)

// WithFailOpen Store 出错时放行（默认是 FailClosed）
func WithFailOpen() Option {
	return func(l *Limiter) { l.failOpen = true }
}

// WithFailClosed Store 出错时拒绝（默认行为）
func WithFailClosed() Option {
	return func(l *Limiter) { l.failOpen = false }
}

// WithCleanupInterval 启动后台定期 Cleanup；interval <= 0 时使用 1 分钟
func WithCleanupInterval(interval time.Duration) Option {
	return func(l *Limiter) { l.cleanupInterval = interval }
}

// Limiter 业务限流器 —— 基于配置的高层封装
type Limiter struct {
	config   *RateLimitConfig
	store    Store
	failOpen bool // false = FailClosed（默认）

	cleanupInterval time.Duration
	cleanupMu       sync.Mutex
	cleanupStop     chan struct{}
	cleanupDone     chan struct{}
	cleanupRunning  bool
}

// NewLimiter 根据配置创建限流器；store 为 nil 时使用默认内存实现。
// 会 Clone 配置快照，外部后续修改不影响已创建的 Limiter。
func NewLimiter(config *RateLimitConfig, store Store, opts ...Option) *Limiter {
	if config == nil {
		config = DefaultConfig()
	} else {
		config = config.Clone()
	}
	if store == nil {
		store = NewSlidingWindowStore(config.maxWindow(), 10)
	}
	l := &Limiter{config: config, store: store}
	for _, opt := range opts {
		opt(l)
	}
	if l.cleanupInterval > 0 {
		l.StartCleanup(l.cleanupInterval)
	}
	return l
}

// NewSimpleLimiter 最常用的便捷 API：创建一个「limit 次/窗口」的简单限流器
func NewSimpleLimiter(limit, window int64, opts ...Option) *Limiter {
	cfg := &RateLimitConfig{
		Enabled: true,
		Global:  &RateLimitItem{Limit: limit, Window: window},
	}
	return NewLimiter(cfg, NewSlidingWindowStore(window, 10), opts...)
}

func (l *Limiter) onStoreError() CheckResult {
	if l.failOpen {
		return CheckResult{}
	}
	return CheckResult{Limited: true, Level: "error"}
}

func fromCount(r CountResult, level string) CheckResult {
	out := CheckResult{
		Limited:    r.Limited,
		ResetAfter: r.ResetAfter,
		Current:    r.Current,
		Remaining:  r.Remaining,
	}
	if r.Limited && level != "" {
		out.Level = level
	}
	return out
}

// Allow 简单判定（不区分 key，使用全局规则）
func (l *Limiter) Allow() CheckResult {
	if !l.config.Enabled {
		return CheckResult{}
	}
	return l.CheckGlobal()
}

// AllowN 自定义 key + 自定义 limit/window 的低层便捷 API
func (l *Limiter) AllowN(key string, limit, window int64) CheckResult {
	if !l.config.Enabled {
		return CheckResult{}
	}
	r, err := l.store.Incr(key, limit, window)
	if err != nil {
		return l.onStoreError()
	}
	return fromCount(r, "")
}

// PeekN 只读查询自定义 key 的限流状态，不消耗配额
func (l *Limiter) PeekN(key string, limit, window int64) CheckResult {
	if !l.config.Enabled {
		return CheckResult{}
	}
	r, err := l.store.Peek(key, limit, window)
	if err != nil {
		return l.onStoreError()
	}
	return fromCount(r, "")
}

// Remaining 返回自定义 key 在窗口内剩余可请求次数；出错或禁用时返回 -1
func (l *Limiter) Remaining(key string, limit, window int64) int64 {
	if !l.config.Enabled {
		return -1
	}
	r, err := l.store.Peek(key, limit, window)
	if err != nil {
		return -1
	}
	return r.Remaining
}

// CheckGlobal 检查全局限流（会 Incr）
func (l *Limiter) CheckGlobal() CheckResult {
	if !l.config.Enabled || l.config.Global == nil {
		return CheckResult{}
	}
	r, err := l.store.Incr("__global__", l.config.Global.Limit, l.config.Global.Window)
	if err != nil {
		return l.onStoreError()
	}
	return fromCount(r, "global")
}

// PeekGlobal 只读查询全局限流状态
func (l *Limiter) PeekGlobal() CheckResult {
	if !l.config.Enabled || l.config.Global == nil {
		return CheckResult{}
	}
	r, err := l.store.Peek("__global__", l.config.Global.Limit, l.config.Global.Window)
	if err != nil {
		return l.onStoreError()
	}
	return fromCount(r, "global")
}

// CheckIP 检查 IP 限流（会 Incr）
func (l *Limiter) CheckIP(ip string) CheckResult {
	if !l.config.Enabled || l.config.PerIP == nil {
		return CheckResult{}
	}
	r, err := l.store.Incr("ip:"+ip, l.config.PerIP.Limit, l.config.PerIP.Window)
	if err != nil {
		return l.onStoreError()
	}
	return fromCount(r, "ip")
}

// PeekIP 只读查询 IP 限流状态
func (l *Limiter) PeekIP(ip string) CheckResult {
	if !l.config.Enabled || l.config.PerIP == nil {
		return CheckResult{}
	}
	r, err := l.store.Peek("ip:"+ip, l.config.PerIP.Limit, l.config.PerIP.Window)
	if err != nil {
		return l.onStoreError()
	}
	return fromCount(r, "ip")
}

// CheckRoute 检查路由限流（会 Incr）
func (l *Limiter) CheckRoute(route string) CheckResult {
	if !l.config.Enabled {
		return CheckResult{}
	}
	rule := l.config.GetRouteRule(route)
	if rule == nil {
		return CheckResult{}
	}
	r, err := l.store.Incr("route:"+route, rule.Limit, rule.Window)
	if err != nil {
		return l.onStoreError()
	}
	return fromCount(r, "route")
}

// PeekRoute 只读查询路由限流状态
func (l *Limiter) PeekRoute(route string) CheckResult {
	if !l.config.Enabled {
		return CheckResult{}
	}
	rule := l.config.GetRouteRule(route)
	if rule == nil {
		return CheckResult{}
	}
	r, err := l.store.Peek("route:"+route, rule.Limit, rule.Window)
	if err != nil {
		return l.onStoreError()
	}
	return fromCount(r, "route")
}

// Check 综合检查：全局 → IP → 路由，任一超限即返回。
// 注意：上级规则在检查时就会 Incr；若后续级别拒绝，已消耗的上级配额不会回滚。
func (l *Limiter) Check(route, ip string) CheckResult {
	if !l.config.Enabled {
		return CheckResult{}
	}
	if r := l.CheckGlobal(); r.Limited {
		return r
	}
	if r := l.CheckIP(ip); r.Limited {
		return r
	}
	if r := l.CheckRoute(route); r.Limited {
		return r
	}
	return CheckResult{}
}

// Peek 综合只读检查：全局 → IP → 路由，不消耗任何配额。
func (l *Limiter) Peek(route, ip string) CheckResult {
	if !l.config.Enabled {
		return CheckResult{}
	}
	if r := l.PeekGlobal(); r.Limited {
		return r
	}
	if r := l.PeekIP(ip); r.Limited {
		return r
	}
	if r := l.PeekRoute(route); r.Limited {
		return r
	}
	return CheckResult{}
}

// StartCleanup 启动后台定期调用 store.Cleanup；重复调用无效。
// interval <= 0 时使用 1 分钟。可用 StopCleanup 停止。
func (l *Limiter) StartCleanup(interval time.Duration) {
	if interval <= 0 {
		interval = time.Minute
	}
	l.cleanupMu.Lock()
	defer l.cleanupMu.Unlock()
	if l.cleanupRunning {
		return
	}
	stop := make(chan struct{})
	done := make(chan struct{})
	l.cleanupStop = stop
	l.cleanupDone = done
	l.cleanupRunning = true
	go func() {
		defer close(done)
		t := time.NewTicker(interval)
		defer t.Stop()
		for {
			select {
			case <-stop:
				return
			case <-t.C:
				_ = l.store.Cleanup()
			}
		}
	}()
}

// StopCleanup 停止后台 Cleanup，并等待 goroutine 退出
func (l *Limiter) StopCleanup() {
	l.cleanupMu.Lock()
	if !l.cleanupRunning {
		l.cleanupMu.Unlock()
		return
	}
	close(l.cleanupStop)
	done := l.cleanupDone
	l.cleanupRunning = false
	l.cleanupMu.Unlock()
	<-done
}

// Config 返回配置深拷贝，修改返回值不影响 Limiter 内部状态
func (l *Limiter) Config() *RateLimitConfig { return l.config.Clone() }

// Store 返回当前存储实现
func (l *Limiter) Store() Store { return l.store }
