package fn

import "testing"

func BenchmarkAnd(b *testing.B) {
	isPositive := func(v int) bool { return v > 0 }
	isEven := func(v int) bool { return v%2 == 0 }
	fn := And(isPositive, isEven)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(42)
	}
}

func BenchmarkOr(b *testing.B) {
	isNegative := func(v int) bool { return v < 0 }
	isZero := func(v int) bool { return v == 0 }
	fn := Or(isNegative, isZero)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(0)
	}
}

func BenchmarkNot(b *testing.B) {
	isZero := func(v int) bool { return v == 0 }
	fn := Not(isZero)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(42)
	}
}
