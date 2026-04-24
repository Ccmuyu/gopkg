package maps

import (
	"testing"
)

func BenchmarkKeys(b *testing.B) {
	m := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Keys(m)
	}
}

func BenchmarkValues(b *testing.B) {
	m := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Values(m)
	}
}

func BenchmarkHasKey(b *testing.B) {
	m := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HasKey(m, 500)
	}
}

func BenchmarkMerge(b *testing.B) {
	m1 := make(map[int]int, 1000)
	m2 := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		m1[i] = i
		m2[i+1000] = i + 1000
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Merge(m1, m2)
	}
}