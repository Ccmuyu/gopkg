package set

import (
	"testing"
)

func BenchmarkSetAdd(b *testing.B) {
	s := New[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
}

func BenchmarkSetContains(b *testing.B) {
	s := New[int]()
	for i := 0; i < 1000; i++ {
		s.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Contains(500)
	}
}

func BenchmarkSetUnion(b *testing.B) {
	s1 := New[int]()
	s2 := New[int]()
	for i := 0; i < 500; i++ {
		s1.Add(i)
	}
	for i := 500; i < 1000; i++ {
		s2.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s1.Union(s2)
	}
}

func BenchmarkSetIntersection(b *testing.B) {
	s1 := New[int]()
	s2 := New[int]()
	for i := 0; i < 1000; i++ {
		s1.Add(i)
	}
	for i := 500; i < 1500; i++ {
		s2.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s1.Intersection(s2)
	}
}

func BenchmarkSetToSlice(b *testing.B) {
	s := New[int]()
	for i := 0; i < 1000; i++ {
		s.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.ToSlice()
	}
}
