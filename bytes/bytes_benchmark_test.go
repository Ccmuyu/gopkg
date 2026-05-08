package bytes

import (
	"testing"
)

func BenchmarkHexEncode(b *testing.B) {
	data := make([]byte, 1024)
	for i := range data {
		data[i] = byte(i % 256)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HexEncode(data)
	}
}

func BenchmarkHexDecode(b *testing.B) {
	data := make([]byte, 1024)
	for i := range data {
		data[i] = byte(i % 256)
	}
	s := HexEncode(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HexDecode(s)
	}
}

func BenchmarkBase64Encode(b *testing.B) {
	data := make([]byte, 1024)
	for i := range data {
		data[i] = byte(i % 256)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Base64Encode(data)
	}
}

func BenchmarkBase64EncodeString(b *testing.B) {
	s := "hello world example string for base64 encoding"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Base64EncodeString(s)
	}
}

func BenchmarkBase64Decode(b *testing.B) {
	s := Base64EncodeString("hello world example string for base64 encoding")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Base64Decode(s)
	}
}

func BenchmarkBase64URLEncodeString(b *testing.B) {
	s := "hello?world=foo&bar=baz"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Base64URLEncodeString(s)
	}
}

func BenchmarkBase64URLDecode(b *testing.B) {
	s := Base64URLEncodeString("hello?world=foo&bar=baz")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Base64URLDecode(s)
	}
}

func BenchmarkHumanReadable(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HumanReadable(1048576)
	}
}