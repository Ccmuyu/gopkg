package file

import (
	"os"
	"testing"
)

func BenchmarkRead(b *testing.B) {
	path := "bench_read.tmp"
	WriteString(path, "benchmark content", 0644)
	defer os.Remove(path)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Read(path)
	}
}

func BenchmarkWrite(b *testing.B) {
	path := "bench_write.tmp"
	data := []byte("benchmark content")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Write(path, data, 0644)
	}
	os.Remove(path)
}

func BenchmarkExists(b *testing.B) {
	path := "bench_exists.tmp"
	WriteString(path, "exists", 0644)
	defer os.Remove(path)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Exists(path)
	}
}

func BenchmarkCopy(b *testing.B) {
	src := "bench_copy_src.tmp"
	dst := "bench_copy_dst.tmp"
	WriteString(src, "benchmark copy content "+string(rune(b.N)), 0644)
	defer os.Remove(src)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Copy(src, dst)
	}
	os.Remove(dst)
}
