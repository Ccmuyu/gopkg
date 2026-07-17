// Package csrf 提供基于双重提交 Cookie 的 CSRF 防护中间件。
package csrf

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultCookieName = "csrf_token"
	defaultHeaderName = "X-CSRF-Token"
	defaultFieldName  = "csrf_token"
	defaultTokenBytes = 32
)

var (
	// ErrMissingToken 请求中缺少 CSRF 令牌。
	ErrMissingToken = errors.New("csrf: missing token")
	// ErrInvalidToken Cookie 与提交令牌不匹配。
	ErrInvalidToken = errors.New("csrf: invalid token")
	// ErrNoCookie 请求尚无 CSRF Cookie。
	ErrNoCookie = errors.New("csrf: no cookie")
)

// Config CSRF 配置。
type Config struct {
	// CookieName Cookie 名，默认 csrf_token。
	CookieName string
	// HeaderName 请求头名，默认 X-CSRF-Token。
	HeaderName string
	// FieldName 表单字段名，默认 csrf_token。
	FieldName string
	// CookiePath Cookie Path，默认 "/"。
	CookiePath string
	// CookieDomain Cookie Domain；空表示宿主默认。
	CookieDomain string
	// Secure 是否仅 HTTPS 发送 Cookie。
	Secure bool
	// HTTPOnly 是否禁止 JS 读取 Cookie。双重提交场景通常为 false。
	HTTPOnly bool
	// SameSite Cookie SameSite；零值使用 SameSiteLaxMode。
	SameSite http.SameSite
	// MaxAge Cookie 有效期（秒）；0 表示会话 Cookie，<0 删除。
	MaxAge int
	// TokenBytes 令牌随机字节数，默认 32。
	TokenBytes int
	// TrustedOrigins 额外信任的 Origin（完整源，如 https://app.example.com）。
	// 非空时，对不安全方法还会校验 Origin / Referer。
	TrustedOrigins []string
	// ErrorHandler 校验失败时的处理；默认 403。
	ErrorHandler func(http.ResponseWriter, *http.Request, error)
}

// Protector CSRF 防护器。
type Protector struct {
	cookieName string
	headerName string
	fieldName  string
	cookiePath string
	domain     string
	secure     bool
	httpOnly   bool
	sameSite   http.SameSite
	maxAge     int
	tokenBytes int
	origins    map[string]struct{}
	onError    func(http.ResponseWriter, *http.Request, error)
}

// New 创建 CSRF 防护器。
func New(cfg Config) *Protector {
	p := &Protector{
		cookieName: cfg.CookieName,
		headerName: cfg.HeaderName,
		fieldName:  cfg.FieldName,
		cookiePath: cfg.CookiePath,
		domain:     cfg.CookieDomain,
		secure:     cfg.Secure,
		httpOnly:   cfg.HTTPOnly,
		sameSite:   cfg.SameSite,
		maxAge:     cfg.MaxAge,
		tokenBytes: cfg.TokenBytes,
		onError:    cfg.ErrorHandler,
	}
	if p.cookieName == "" {
		p.cookieName = defaultCookieName
	}
	if p.headerName == "" {
		p.headerName = defaultHeaderName
	}
	if p.fieldName == "" {
		p.fieldName = defaultFieldName
	}
	if p.cookiePath == "" {
		p.cookiePath = "/"
	}
	if p.sameSite == 0 {
		p.sameSite = http.SameSiteLaxMode
	}
	if p.tokenBytes <= 0 {
		p.tokenBytes = defaultTokenBytes
	}
	if p.onError == nil {
		p.onError = defaultErrorHandler
	}
	if len(cfg.TrustedOrigins) > 0 {
		p.origins = make(map[string]struct{}, len(cfg.TrustedOrigins))
		for _, o := range cfg.TrustedOrigins {
			o = strings.TrimSpace(o)
			if o != "" {
				p.origins[o] = struct{}{}
			}
		}
	}
	return p
}

// Generate 生成新的 CSRF 令牌。
func (p *Protector) Generate() (string, error) {
	b := make([]byte, p.tokenBytes)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// Token 返回请求中的 Cookie 令牌；若不存在则生成新令牌。
// 调用方应在响应中通过 SetCookie 下发。
func (p *Protector) Token(r *http.Request) (string, error) {
	if c, err := r.Cookie(p.cookieName); err == nil && c.Value != "" {
		return c.Value, nil
	}
	return p.Generate()
}

// SetCookie 将令牌写入响应 Cookie。
func (p *Protector) SetCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     p.cookieName,
		Value:    token,
		Path:     p.cookiePath,
		Domain:   p.domain,
		MaxAge:   p.maxAge,
		Secure:   p.secure,
		HttpOnly: p.httpOnly,
		SameSite: p.sameSite,
	})
}

// Valid 校验请求是否携带与 Cookie 一致的 CSRF 令牌。
func (p *Protector) Valid(r *http.Request) error {
	cookie, err := r.Cookie(p.cookieName)
	if err != nil || cookie.Value == "" {
		return ErrNoCookie
	}
	submitted := p.submittedToken(r)
	if submitted == "" {
		return ErrMissingToken
	}
	if subtle.ConstantTimeCompare([]byte(cookie.Value), []byte(submitted)) != 1 {
		return ErrInvalidToken
	}
	if len(p.origins) > 0 {
		if err := p.checkOrigin(r); err != nil {
			return err
		}
	}
	return nil
}

// Middleware 对不安全方法校验 CSRF；安全方法仅确保 Cookie 存在。
func (p *Protector) Middleware() func(http.Handler) http.Handler {
	return p.Handler
}

// Handler 返回包装后的 http.Handler。
func (p *Protector) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := p.Token(r)
		if err != nil {
			p.onError(w, r, err)
			return
		}
		// 每次响应都刷新 Cookie，便于前端读取（HTTPOnly=false 时）。
		p.SetCookie(w, token)

		if isSafeMethod(r.Method) {
			next.ServeHTTP(w, r)
			return
		}
		if err := p.Valid(r); err != nil {
			p.onError(w, r, err)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (p *Protector) submittedToken(r *http.Request) string {
	if v := r.Header.Get(p.headerName); v != "" {
		return v
	}
	if err := r.ParseForm(); err == nil {
		if v := r.PostForm.Get(p.fieldName); v != "" {
			return v
		}
		if v := r.Form.Get(p.fieldName); v != "" {
			return v
		}
	}
	return ""
}

func (p *Protector) checkOrigin(r *http.Request) error {
	origin := r.Header.Get("Origin")
	if origin == "" {
		if ref := r.Header.Get("Referer"); ref != "" {
			u, err := url.Parse(ref)
			if err != nil {
				return ErrInvalidToken
			}
			origin = u.Scheme + "://" + u.Host
		}
	}
	if origin == "" {
		return ErrInvalidToken
	}
	if _, ok := p.origins[origin]; ok {
		return nil
	}
	return ErrInvalidToken
}

func isSafeMethod(method string) bool {
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodTrace:
		return true
	default:
		return false
	}
}

func defaultErrorHandler(w http.ResponseWriter, _ *http.Request, _ error) {
	http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
}
