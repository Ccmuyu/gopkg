package coalesce

import "testing"

func BenchmarkCoalesce(b *testing.B) {
	vals := []int{0, 0, 0, 42, 100}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Coalesce(vals...)
	}
}

func BenchmarkCoalesceString(b *testing.B) {
	vals := []string{"", "", "", "found", "ignored"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Coalesce(vals...)
	}
}
