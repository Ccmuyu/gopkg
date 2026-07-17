package rate

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestMiddlewareAllows(t *testing.T) {
	cfg := NewConfig(
		&RateLimitItem{Limit: 10, Window: 60},
		&RateLimitItem{Limit: 10, Window: 60},
		RateLimitRouteItem{Route: "/ok", Limit: 10, Window: 60},
	)
	l := NewLimiter(cfg, nil)

	called := false
	h := l.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/ok", nil)
	req.RemoteAddr = "10.0.0.1:1234"
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	AssertTrue(t, called)
	AssertEqual(t, rr.Code, http.StatusOK)
}

func TestMiddlewareLimits(t *testing.T) {
	cfg := NewConfig(
		nil,
		nil,
		RateLimitRouteItem{Route: "/api/sync", Limit: 1, Window: 60},
	)
	l := NewLimiter(cfg, nil)

	h := l.Middleware()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/api/sync", nil)
	req.RemoteAddr = "10.0.0.1:1234"

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	AssertEqual(t, rr.Code, http.StatusOK)

	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	AssertEqual(t, rr.Code, http.StatusTooManyRequests)
	AssertEqual(t, rr.Header().Get("X-RateLimit-Level"), "route")
}

func TestMiddlewareCustomOnLimited(t *testing.T) {
	cfg := NewConfig(nil, nil, RateLimitRouteItem{Route: "/x", Limit: 0, Window: 60})
	l := NewLimiter(cfg, nil)

	h := l.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("should not reach next")
	}), WithOnLimited(func(w http.ResponseWriter, r *http.Request, res CheckResult) {
		w.WriteHeader(http.StatusTeapot)
		_, _ = w.Write([]byte(res.Level))
	}))

	// Limit 0 → Incr 后 current=1 > 0 → limited
	// Wait, Limit 0: store returns Limited true without incr for limit<=0
	req := httptest.NewRequest(http.MethodGet, "/x", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	AssertEqual(t, rr.Code, http.StatusTeapot)
}

func TestDefaultClientIP(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Forwarded-For", "1.2.3.4, 5.6.7.8")
	req.RemoteAddr = "8.8.8.8:9999"
	AssertEqual(t, DefaultClientIP(req), "8.8.8.8")
}

func TestForwardedClientIP(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Forwarded-For", "1.2.3.4, 5.6.7.8")
	AssertEqual(t, ForwardedClientIP(req), "1.2.3.4")

	req = httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Real-IP", "9.9.9.9")
	AssertEqual(t, ForwardedClientIP(req), "9.9.9.9")

	req = httptest.NewRequest(http.MethodGet, "/", nil)
	req.RemoteAddr = "8.8.8.8:9999"
	AssertEqual(t, ForwardedClientIP(req), "8.8.8.8")
}

func TestMiddlewareCustomIPAndRoute(t *testing.T) {
	cfg := NewConfig(nil, &RateLimitItem{Limit: 1, Window: 60})
	l := NewLimiter(cfg, nil)

	h := l.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}),
		WithClientIP(func(r *http.Request) string { return "fixed-ip" }),
		WithRouteKey(func(r *http.Request) string { return "/ignored" }),
	)

	req := httptest.NewRequest(http.MethodGet, "/any", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	AssertEqual(t, rr.Code, http.StatusOK)

	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	AssertEqual(t, rr.Code, http.StatusTooManyRequests)
	AssertEqual(t, rr.Header().Get("X-RateLimit-Level"), "ip")
}
