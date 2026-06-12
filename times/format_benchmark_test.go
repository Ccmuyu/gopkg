package times

import (
	"testing"
	"time"
)

func BenchmarkFormat(b *testing.B) {
	tm := time.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Format(tm)
	}
}

func BenchmarkBeginOfDay(b *testing.B) {
	tm := time.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BeginOfDay(tm)
	}
}

func BenchmarkEndOfDay(b *testing.B) {
	tm := time.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EndOfDay(tm)
	}
}

func BenchmarkBeginOfMonth(b *testing.B) {
	tm := time.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BeginOfMonth(tm)
	}
}

func BenchmarkDaysBetween(b *testing.B) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DaysBetween(start, end)
	}
}

func BenchmarkIsWeekend(b *testing.B) {
	tm := time.Date(2024, 9, 21, 0, 0, 0, 0, time.UTC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsWeekend(tm)
	}
}
