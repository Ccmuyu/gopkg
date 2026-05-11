package maths

import "testing"

func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Min(12345, 67890)
	}
}

func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Max(12345, 67890)
	}
}

func BenchmarkClamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Clamp(50, 1, 100)
	}
}

func BenchmarkAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Abs(-12345)
	}
}

func BenchmarkSum(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		Sum(nums...)
	}
}

func BenchmarkAverage(b *testing.B) {
	nums := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	for i := 0; i < b.N; i++ {
		Average(nums...)
	}
}

func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pow(2, 20)
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(9973)
	}
}

func BenchmarkGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCD(123456, 789012)
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(30)
	}
}
