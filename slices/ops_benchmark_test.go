package slices

import (
	"testing"
)

func BenchmarkDedup(b *testing.B) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = i % 100
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Dedup(arr)
	}
}

func BenchmarkFilter(b *testing.B) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(arr, func(v int) bool {
			return v%2 == 0
		})
	}
}

func BenchmarkMap(b *testing.B) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Map(arr, func(v int) int {
			return v * 2
		})
	}
}

func BenchmarkContains(b *testing.B) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Contains(arr, 9999)
	}
}