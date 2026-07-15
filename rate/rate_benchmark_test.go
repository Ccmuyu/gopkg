package rate

import (
	"testing"
)

func BenchmarkAllow(b *testing.B) {
	l := NewSimpleLimiter(1000000, 60)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = l.Allow()
	}
}

func BenchmarkCheck(b *testing.B) {
	cfg := NewConfig(
		&RateLimitItem{Limit: 1000000, Window: 60},
		&RateLimitItem{Limit: 1000, Window: 60},
		RateLimitRouteItem{Route: "/api/test", Limit: 100, Window: 60},
	)
	l := NewLimiter(cfg, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = l.Check("/api/test", "127.0.0.1")
	}
}

func BenchmarkStoreIncr(b *testing.B) {
	s := NewSlidingWindowStore(60, 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = s.Incr("k", 1000000, 60)
	}
}
