package rand

import "testing"

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(16)
	}
}

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int(0, 1000)
	}
}

func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(32)
	}
}

func BenchmarkChoice(b *testing.B) {
	slice := []int{10, 20, 30, 40, 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Choice(slice)
	}
}

func BenchmarkShuffle(b *testing.B) {
	slice := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Shuffle(slice)
	}
}
