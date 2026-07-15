package rate_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/Ccmuyu/gopkg/rate"
)

func ExampleNewSimpleLimiter() {
	// 最简单的用法：全局 3 次 / 60s
	l := rate.NewSimpleLimiter(3, 60)

	for i := 0; i < 4; i++ {
		r := l.Allow()
		if r.Limited {
			fmt.Println("limited")
			break
		} else {
			fmt.Println("pass")
		}
	}
	// Output:
	// pass
	// pass
	// pass
	// limited
}

func ExampleLimiter_Check() {
	cfg := rate.NewConfig(
		&rate.RateLimitItem{Limit: 100, Window: 60},
		&rate.RateLimitItem{Limit: 3, Window: 60},
		rate.RateLimitRouteItem{Route: "/api/creator/sync", Limit: 5, Window: 60},
	)
	l := rate.NewLimiter(cfg, nil)

	r := l.Check("/api/creator/sync", "10.0.0.1")
	if r.Limited {
		fmt.Printf("limited by %s, retry after %d s\n", r.Level, r.ResetAfter)
	} else {
		fmt.Println("pass")
	}
	// Output: pass
}

func ExampleNewLimiter() {
	cfg := &rate.RateLimitConfig{
		Enabled: true,
		Global:  &rate.RateLimitItem{Limit: 5, Window: 60},
	}
	l := rate.NewLimiter(cfg, nil)
	r := l.CheckGlobal()
	if r.Limited {
		fmt.Println("limited")
	} else {
		fmt.Println("pass")
	}
	// Output: pass
}

func ExampleLimiter_Peek() {
	l := rate.NewSimpleLimiter(2, 60)
	_ = l.Allow()
	_ = l.Allow()

	r := l.PeekGlobal()
	fmt.Println("limited:", r.Limited, "remaining:", r.Remaining)
	// Output: limited: true remaining: 0
}

func ExampleLimiter_Middleware() {
	cfg := rate.NewConfig(nil, nil,
		rate.RateLimitRouteItem{Route: "/api/sync", Limit: 1, Window: 60},
	)
	l := rate.NewLimiter(cfg, nil)

	h := l.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	}))

	req := httptest.NewRequest(http.MethodGet, "/api/sync", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	fmt.Println(rr.Code)

	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	fmt.Println(rr.Code)
	// Output:
	// 200
	// 429
}

func ExampleDefaultConfig() {
	cfg := rate.DefaultConfig()
	fmt.Println("enabled:", cfg.Enabled)
	fmt.Println("global limit:", cfg.Global.Limit)
	// Output:
	// enabled: false
	// global limit: 1000
}
