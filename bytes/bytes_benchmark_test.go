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

func BenchmarkHumanReadable(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HumanReadable(1048576)
	}
}