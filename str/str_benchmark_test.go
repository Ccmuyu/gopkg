package str

import (
	"strings"
	"testing"
)

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

func BenchmarkAfter(b *testing.B) {
	s := "user@example.com"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		After(s, "@")
	}
}

func BenchmarkReverse(b *testing.B) {
	s := "hello world example string"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reverse(s)
	}
}

func BenchmarkToCamel(b *testing.B) {
	s := "hello_world_example_string"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToCamel(s)
	}
}

func BenchmarkMask(b *testing.B) {
	s := "13812345678"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Mask(s, 3, '*')
	}
}