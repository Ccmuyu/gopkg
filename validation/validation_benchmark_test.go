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