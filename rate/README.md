# rate

基于滑动窗口计数器的限流器，支持全局、IP、路由三级限流；线程安全，默认内存实现，可接 Redis，并提供 `net/http` 中间件。

## 安装

```go
import "github.com/Ccmuyu/gopkg/rate"
```

## 快速开始

### 简单限流器（全局计数）

```go
// 每 60 秒最多 5 次请求
l := rate.NewSimpleLimiter(5, 60)
if r := l.Allow(); r.Limited {
    fmt.Printf("请求过于频繁，%d 秒后重试\n", r.ResetAfter)
}
```

### 三级组合限流（推荐）

```go
cfg := rate.NewConfig(
    &rate.RateLimitItem{Limit: 1000, Window: 60}, // 全局限流；传 nil 表示不启用该级
    &rate.RateLimitItem{Limit: 100,  Window: 60}, // 单 IP 限流
    rate.RateLimitRouteItem{Route: "/api/creator/sync",           Limit: 5,  Window: 60},
    rate.RateLimitRouteItem{Route: "/api/creator/parse/shortlink", Limit: 30, Window: 60},
    rate.RateLimitRouteItem{Route: "/api/creator/query",           Limit: 200, Window: 60},
)
l := rate.NewLimiter(cfg, nil)

r := l.Check("/api/creator/sync", clientIP)
if r.Limited {
    fmt.Printf("被 %s 规则限流，%d 秒后重试\n", r.Level, r.ResetAfter)
}
```

仅路由限流时，将 global / perIP 传 `nil` 即可（不会继承默认限额）：

```go
cfg := rate.NewConfig(nil, nil,
    rate.RateLimitRouteItem{Route: "/api/sync", Limit: 5, Window: 60},
)
```

### 只读查询（不消耗配额）

```go
if l.Peek("/api/sync", clientIP).Limited {
    // 已达上限，可提前拒绝或降级，不占用计数
}
left := l.Remaining("user:"+id, 100, 60)
```

### 使用默认配置（默认禁用，按需开启）

```go
cfg := rate.DefaultConfig()
cfg.Enabled = true
l := rate.NewLimiter(cfg, nil)
```

### 自定义 key 限流

```go
l := rate.NewSimpleLimiter(100, 60)
r := l.AllowN("user:"+userID, 10, 60)
```

### 失败策略与后台清理

```go
// 默认 FailClosed：Store 出错时拒绝（Level == "error"）
l := rate.NewLimiter(cfg, store)

// 需要可用性优先时改为 FailOpen
l := rate.NewLimiter(cfg, store, rate.WithFailOpen())

// 启动后台定期清理空闲 key（进程退出前 StopCleanup）
l := rate.NewLimiter(cfg, nil, rate.WithCleanupInterval(time.Minute))
defer l.StopCleanup()
```

### HTTP 中间件

```go
mux := http.NewServeMux()
mux.HandleFunc("/api/sync", handleSync)

h := l.Middleware(
    rate.WithClientIP(rate.DefaultClientIP), // 默认已启用，仅信任 RemoteAddr
)(mux)
http.ListenAndServe(":8080", h)
```

被限流时默认返回 **429** 并设置 `X-RateLimit-Level`；仅当等待时间大于 0 时设置 `Retry-After`。

服务位于可信反向代理之后，且代理会覆盖客户端传入的转发头时，可显式启用：

```go
h := l.Middleware(
    rate.WithClientIP(rate.ForwardedClientIP),
)(mux)
```

### Redis Store（无第三方依赖）

通过 `RedisEvaler` 适配任意 Redis 客户端（如 go-redis）：

```go
type redisEvaler struct{ rdb *redis.Client }

func (e redisEvaler) Eval(script string, keys []string, args ...any) (any, error) {
    return e.rdb.Eval(context.Background(), script, keys, args...).Result()
}

store := rate.NewRedisStore(redisEvaler{rdb}, "myapp:rate:")
l := rate.NewLimiter(cfg, store)
```

## 行为说明

- `Check` 按 **全局 → IP → 路由** 依次检查；每一级在检查时就会 `Incr`。
- 若后续级别拒绝，**已消耗的上级配额不会回滚**（被拒请求仍计入上级计数）。
- 已达到某一级限额后，该级后续被拒请求**不会继续增加计数**，窗口到期后可正常恢复。
- 相同业务 key 的不同 `window` 使用独立计数状态，切换窗口不会删除另一窗口的历史。
- 需要「先看后扣」时用 `Peek` / `PeekN` / `Remaining`（不消耗配额）。
- `NewLimiter` 会 **Clone 配置快照**；`Config()` / `GetRouteRule` 返回副本，修改不影响内部状态。
- Store 出错时默认 **FailClosed**（拒绝）；可用 `WithFailOpen` 改为放行。
- 内存 Store 按时间槽近似计算并采用保守边界，最多多限制一个槽宽；Redis Store 按秒级 ZSET 精确计算。
- 内存 Store 会在请求过程中按分钟摊销清理，也可通过 `Cleanup` 主动删除空闲超过 `max(构造 idleTTL, 2*lastWindow)` 的 key。
- `WithCleanupInterval` 显式启动后台清理；传入非正间隔时使用 1 分钟。

## API 参考

```go
// 限流器构造
func NewLimiter(config *RateLimitConfig, store Store, opts ...Option) *Limiter
func NewSimpleLimiter(limit, window int64, opts ...Option) *Limiter
func WithFailOpen() Option
func WithFailClosed() Option
func WithCleanupInterval(interval time.Duration) Option

// 配置
func DefaultConfig() *RateLimitConfig
func NewConfig(global, perIP *RateLimitItem, routes ...RateLimitRouteItem) *RateLimitConfig // nil = 禁用该级
func (c *RateLimitConfig) Clone() *RateLimitConfig
func (c *RateLimitConfig) GetRouteRule(route string) *RateLimitRouteItem // 返回副本

// 核心检查（消耗配额）
func (l *Limiter) Allow() CheckResult
func (l *Limiter) AllowN(key string, limit, window int64) CheckResult
func (l *Limiter) CheckGlobal() CheckResult
func (l *Limiter) CheckIP(ip string) CheckResult
func (l *Limiter) CheckRoute(route string) CheckResult
func (l *Limiter) Check(route, ip string) CheckResult

// 只读查询（不消耗配额）
func (l *Limiter) PeekN(key string, limit, window int64) CheckResult
func (l *Limiter) PeekGlobal() CheckResult
func (l *Limiter) PeekIP(ip string) CheckResult
func (l *Limiter) PeekRoute(route string) CheckResult
func (l *Limiter) Peek(route, ip string) CheckResult
func (l *Limiter) Remaining(key string, limit, window int64) int64

func (l *Limiter) StartCleanup(interval time.Duration)
func (l *Limiter) StopCleanup()
func (l *Limiter) Config() *RateLimitConfig // 深拷贝

// HTTP 中间件
func (l *Limiter) Middleware(opts ...MiddlewareOption) func(http.Handler) http.Handler
func (l *Limiter) Handler(next http.Handler, opts ...MiddlewareOption) http.Handler
func WithClientIP(fn func(*http.Request) string) MiddlewareOption
func WithRouteKey(fn func(*http.Request) string) MiddlewareOption
func WithOnLimited(fn func(http.ResponseWriter, *http.Request, CheckResult)) MiddlewareOption
func WithStatusCode(code int) MiddlewareOption
func DefaultClientIP(r *http.Request) string   // 仅 RemoteAddr，安全默认值
func ForwardedClientIP(r *http.Request) string // 信任 X-Forwarded-For / X-Real-IP

// 存储接口
type Store interface {
    // 仅在当前计数低于 limit 时增加；达到限额后拒绝且不计数。
    Incr(key string, limit, window int64) (CountResult, error)
    Peek(key string, limit, window int64) (CountResult, error)
    Cleanup() error
}
func NewSlidingWindowStore(window int64, slots int) Store
func NewRedisStore(evaler RedisEvaler, keyPrefix string) Store
type RedisEvaler interface {
    Eval(script string, keys []string, args ...any) (any, error)
}

// 结果结构（值类型）
type CheckResult struct {
    Limited    bool
    ResetAfter int64
    Level      string // "global" | "ip" | "route" | "error"
    Current    int64
    Remaining  int64
}
type CountResult struct {
    Current    int64
    Limited    bool
    ResetAfter int64
    Remaining  int64
}
```
