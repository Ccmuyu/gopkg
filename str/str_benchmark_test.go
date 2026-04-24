package str

import (
	"strings"
	"testing"
)

func AssertSliceEqual[T any](t *testing.T, actual, expected []T) {
	if len(actual) != len(expected) {
		t.Errorf("len mismatch: %d vs %d", len(actual), len(expected))
		return
	}
	for i := range actual {
		if any(actual[i]) != any(expected[i]) {
			t.Errorf("index %d: %v vs %v", i, actual[i], expected[i])
		}
	}
}

func BenchmarkUpper(b *testing.B) {
	s := "hello world example string for benchmarking"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Upper(s)
	}
}

func BenchmarkContains(b *testing.B) {
	s := "hello world example string for benchmarking"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Contains(s, "example")
	}
}

func BenchmarkSplit(b *testing.B) {
	s := "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Split(s, ",")
	}
}

func BenchmarkEllipsis(b *testing.B) {
	s := strings.Repeat("a", 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Ellipsis(s, 10)
	}
}