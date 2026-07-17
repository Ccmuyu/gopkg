// Package cors 提供 CORS（跨域资源共享）中间件，基于标准库 net/http。
package cors

import (
	"net/http"
	"strconv"
	"strings"
)

// Config CORS 配置。
type Config struct {
	// AllowOrigins 允许的源。含 "*" 表示任意源（与 AllowCredentials 互斥）。
	AllowOrigins []string
	// AllowMethods 预检允许的方法；为空时默认 GET、POST、PUT、PATCH、DELETE、HEAD、OPTIONS。
	AllowMethods []string
	// AllowHeaders 预检允许的请求头；为空时回显 Access-Control-Request-Headers。
	AllowHeaders []string
	// ExposeHeaders 允许浏览器读取的响应头。
	ExposeHeaders []string
	// AllowCredentials 是否允许携带 Cookie / Authorization。
	AllowCredentials bool
	// MaxAge 预检结果缓存秒数；<=0 不设置 Access-Control-Max-Age。
	MaxAge int
	// AllowPrivateNetwork 是否允许私有网络访问（Chrome Private Network Access）。
	AllowPrivateNetwork bool
}

// CORS 跨域处理器。
type CORS struct {
	allowAll    bool
	origins     map[string]struct{}
	methods     string
	headers     string
	expose      string
	credentials bool
	maxAge      string
	privateNet  bool
}

// Default 返回常用开发默认：允许任意源、常见方法，不携带凭证。
func Default() *CORS {
	return New(Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet, http.MethodPost, http.MethodPut,
			http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions,
		},
	})
}

// New 根据配置创建 CORS。AllowCredentials 为 true 且 AllowOrigins 含 "*" 时，
// "*" 会被忽略（规范禁止凭证模式下使用通配源）。
func New(cfg Config) *CORS {
	c := &CORS{
		credentials: cfg.AllowCredentials,
		privateNet:  cfg.AllowPrivateNetwork,
	}

	c.origins = make(map[string]struct{}, len(cfg.AllowOrigins))
	for _, o := range cfg.AllowOrigins {
		o = strings.TrimSpace(o)
		if o == "" {
			continue
		}
		if o == "*" {
			if !cfg.AllowCredentials {
				c.allowAll = true
			}
			continue
		}
		c.origins[o] = struct{}{}
	}

	methods := cfg.AllowMethods
	if len(methods) == 0 {
		methods = []string{
			http.MethodGet, http.MethodPost, http.MethodPut,
			http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions,
		}
	}
	c.methods = strings.Join(methods, ", ")

	if len(cfg.AllowHeaders) > 0 {
		c.headers = strings.Join(cfg.AllowHeaders, ", ")
	}
	if len(cfg.ExposeHeaders) > 0 {
		c.expose = strings.Join(cfg.ExposeHeaders, ", ")
	}
	if cfg.MaxAge > 0 {
		c.maxAge = strconv.Itoa(cfg.MaxAge)
	}
	return c
}

// Middleware 返回标准库中间件。
func (c *CORS) Middleware() func(http.Handler) http.Handler {
	return c.Handler
}

// Handler 处理预检并在实际请求上附加 CORS 响应头。
func (c *CORS) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		allowed, value := c.matchOrigin(origin)

		if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
			if allowed {
				c.writeHeaders(w, value, r, true)
				w.WriteHeader(http.StatusNoContent)
				return
			}
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if allowed {
			c.writeHeaders(w, value, r, false)
		}
		next.ServeHTTP(w, r)
	})
}

func (c *CORS) matchOrigin(origin string) (bool, string) {
	if origin == "" {
		return false, ""
	}
	if c.allowAll {
		return true, "*"
	}
	if _, ok := c.origins[origin]; ok {
		return true, origin
	}
	return false, ""
}

func (c *CORS) writeHeaders(w http.ResponseWriter, allowOrigin string, r *http.Request, preflight bool) {
	h := w.Header()
	h.Set("Access-Control-Allow-Origin", allowOrigin)
	if allowOrigin != "*" {
		h.Add("Vary", "Origin")
	}
	if c.credentials {
		h.Set("Access-Control-Allow-Credentials", "true")
	}
	if c.expose != "" {
		h.Set("Access-Control-Expose-Headers", c.expose)
	}
	if !preflight {
		return
	}

	h.Set("Access-Control-Allow-Methods", c.methods)
	if c.headers != "" {
		h.Set("Access-Control-Allow-Headers", c.headers)
	} else if reqHeaders := r.Header.Get("Access-Control-Request-Headers"); reqHeaders != "" {
		h.Set("Access-Control-Allow-Headers", reqHeaders)
		h.Add("Vary", "Access-Control-Request-Headers")
	}
	if c.maxAge != "" {
		h.Set("Access-Control-Max-Age", c.maxAge)
	}
	if c.privateNet && r.Header.Get("Access-Control-Request-Private-Network") == "true" {
		h.Set("Access-Control-Allow-Private-Network", "true")
	}
	h.Add("Vary", "Access-Control-Request-Method")
}
