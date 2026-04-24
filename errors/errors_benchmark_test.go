package errors

import (
	"testing"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New("error message")
	}
}

func BenchmarkWrap(b *testing.B) {
	err := New("original")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Wrap(err, "wrap message")
	}
}

func BenchmarkIs(b *testing.B) {
	sentinel := New("sentinel")
	wrapped := Wrap(sentinel, "wrapped")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Is(wrapped, sentinel)
	}
}