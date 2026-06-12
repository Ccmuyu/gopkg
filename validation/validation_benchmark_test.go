package validation

import (
	"testing"
)

func BenchmarkIsEmail(b *testing.B) {
	s := "test@example.com"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsEmail(s)
	}
}

func BenchmarkIsPhone(b *testing.B) {
	s := "13812345678"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPhone(s)
	}
}

func BenchmarkIsURL(b *testing.B) {
	s := "https://example.com/path/to/resource"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsURL(s)
	}
}

func BenchmarkIsIP(b *testing.B) {
	s := "192.168.1.1"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsIP(s)
	}
}

func BenchmarkIsMatch(b *testing.B) {
	s := "hello123world"
	pattern := `\d+`
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsMatch(s, pattern)
	}
}

func BenchmarkIsUUID(b *testing.B) {
	s := "550e8400-e29b-41d4-a716-446655440000"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsUUID(s)
	}
}

func BenchmarkIsJSON(b *testing.B) {
	s := `{"key":"value","number":42}`
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsJSON(s)
	}
}

func BenchmarkIsCreditCard(b *testing.B) {
	s := "4111111111111111"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsCreditCard(s)
	}
}