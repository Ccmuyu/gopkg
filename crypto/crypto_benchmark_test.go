package crypto

import (
	"testing"
)

func BenchmarkMD5(b *testing.B) {
	s := "hello world example string"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MD5(s)
	}
}

func BenchmarkSHA1(b *testing.B) {
	s := "hello world example string"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SHA1(s)
	}
}

func BenchmarkSHA256(b *testing.B) {
	s := "hello world example string"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SHA256(s)
	}
}