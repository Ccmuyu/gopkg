package slices

import (
	"testing"
)

func BenchmarkDedup(b *testing.B) {
	arr := []int{1, 2, 3, 2, 4, 3, 5, 1, 6, 7, 8, 7, 9, 0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Dedup(arr)
	}
}

func BenchmarkFilter(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(arr, func(v int) bool { return v%2 == 0 })
	}
}

func BenchmarkMap(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Map(arr, func(v int) int { return v * 2 })
	}
}

func BenchmarkContains(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Contains(arr, 999)
	}
}

func BenchmarkSplit(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Split(arr, 10)
	}
}

func BenchmarkGroupBy(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GroupBy(arr, func(v int) int { return v % 10 })
	}
}
