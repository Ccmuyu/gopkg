package rate

import (
	"net"
	"net/http"
	"strconv"
	"strings"
)

// MiddlewareOption 配置 HTTP 中间件行为
type MiddlewareOption func(*middlewareConfig)

type middlewareConfig struct {
	ipFunc     func(*http.Request) string
	routeFunc  func(*http.Request) string
	onLimited  func(http.ResponseWriter, *http.Request, CheckResult)
	statusCode int
}

// WithClientIP 自定义从请求提取客户端 IP
func WithClientIP(fn func(*http.Request) string) MiddlewareOption {
	return func(c *middlewareConfig) { c.ipFunc = fn }
}

// WithRouteKey 自定义从请求提取路由键（默认 r.URL.Path）
func WithRouteKey(fn func(*http.Request) string) MiddlewareOption {
	return func(c *middlewareConfig) { c.routeFunc = fn }
}

// WithOnLimited 自定义被限流时的响应；默认写 429 + Retry-After
func WithOnLimited(fn func(http.ResponseWriter, *http.Request, CheckResult)) MiddlewareOption {
	return func(c *middlewareConfig) { c.onLimited = fn }
}

// WithStatusCode 被限流时的 HTTP 状态码（默认 429）
func WithStatusCode(code int) MiddlewareOption {
	return func(c *middlewareConfig) {
		if code > 0 {
			c.statusCode = code
		}
	}
}

// Middleware 返回标准库中间件：按全局 → IP → 路由检查，通过则调用 next。
func (l *Limiter) Middleware(opts ...MiddlewareOption) func(http.Handler) http.Handler {
	cfg := &middlewareConfig{
		ipFunc:     DefaultClientIP,
		routeFunc:  defaultRouteKey,
		statusCode: http.StatusTooManyRequests,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	if cfg.onLimited == nil {
		status := cfg.statusCode
		cfg.onLimited = func(w http.ResponseWriter, _ *http.Request, r CheckResult) {
			if r.ResetAfter > 0 {
				w.Header().Set("Retry-After", strconv.FormatInt(r.ResetAfter, 10))
			}
			if r.Level != "" {
				w.Header().Set("X-RateLimit-Level", r.Level)
			}
			http.Error(w, http.StatusText(status), status)
		}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			route := cfg.routeFunc(req)
			ip := cfg.ipFunc(req)
			if r := l.Check(route, ip); r.Limited {
				cfg.onLimited(w, req, r)
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}

// Handler 等价于 Middleware()(next)
func (l *Limiter) Handler(next http.Handler, opts ...MiddlewareOption) http.Handler {
	return l.Middleware(opts...)(next)
}

func defaultRouteKey(r *http.Request) string {
	if r.URL == nil {
		return ""
	}
	return r.URL.Path
}

// DefaultClientIP 从 X-Forwarded-For / X-Real-IP / RemoteAddr 提取客户端 IP
func DefaultClientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		if ip := strings.TrimSpace(parts[0]); ip != "" {
			return ip
		}
	}
	if xri := strings.TrimSpace(r.Header.Get("X-Real-IP")); xri != "" {
		return xri
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
