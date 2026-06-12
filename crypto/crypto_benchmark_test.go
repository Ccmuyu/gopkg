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
	for i := 0; i < b.N; i++ {
		SHA256("hello world example string")
	}
}

func BenchmarkSHA512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SHA512("hello world example string")
	}
}

func BenchmarkCRC32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CRC32([]byte("hello world example string"))
	}
}

func BenchmarkRandomToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomToken(32)
	}
}