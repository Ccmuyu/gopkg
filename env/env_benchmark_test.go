package env

import (
	"os"
	"testing"
)

func BenchmarkGet(b *testing.B) {
	os.Setenv("BENCH_KEY", "value")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Get("BENCH_KEY", "default")
	}
}

func BenchmarkEnvironMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EnvironMap()
	}
}

func BenchmarkHas(b *testing.B) {
	os.Setenv("BENCH_KEY2", "value")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Has("BENCH_KEY2")
	}
}
