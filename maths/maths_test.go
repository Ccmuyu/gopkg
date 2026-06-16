package maths

import (
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestMin(t *testing.T) {
	AssertEqual(t, Min(1, 2), 1)
	AssertEqual(t, Min(2, 1), 1)
	AssertEqual(t, Min(1, 1), 1)
	AssertEqual(t, Min(3.14, 2.71), 2.71)
	AssertEqual(t, Min("a", "b"), "a")
}

func TestMax(t *testing.T) {
	AssertEqual(t, Max(1, 2), 2)
	AssertEqual(t, Max(2, 1), 2)
	AssertEqual(t, Max(1, 1), 1)
	AssertEqual(t, Max(3.14, 2.71), 3.14)
	AssertEqual(t, Max("a", "b"), "b")
}

func TestClamp(t *testing.T) {
	AssertEqual(t, Clamp(5, 1, 10), 5)
	AssertEqual(t, Clamp(0, 1, 10), 1)
	AssertEqual(t, Clamp(15, 1, 10), 10)
	AssertEqual(t, Clamp(1, 1, 10), 1)
	AssertEqual(t, Clamp(10, 1, 10), 10)
	AssertEqual(t, Clamp(3.5, 1.0, 5.0), 3.5)
}

func TestAbs(t *testing.T) {
	AssertEqual(t, Abs(5), 5)
	AssertEqual(t, Abs(-5), 5)
	AssertEqual(t, Abs(0), 0)
	AssertEqual(t, Abs(-3.14), 3.14)
	AssertEqual(t, Abs(3.14), 3.14)
}

func TestAbsInt8(t *testing.T) {
	var x int8 = -7
	AssertEqual(t, Abs(x), int8(7))
}

func TestAbsInt8Min(t *testing.T) {
	var x int8 = -128
	AssertEqual(t, Abs(x), int8(-128))
}

func TestSum(t *testing.T) {
	AssertEqual(t, Sum(1, 2, 3, 4, 5), 15)
	AssertEqual(t, Sum(1), 1)
	AssertEqual(t, Sum[int](), 0)
	AssertEqual(t, Sum(1.5, 2.5, 3.0), 7.0)
}

func TestAverage(t *testing.T) {
	AssertEqual(t, Average(1, 2, 3, 4, 5), 3.0)
	AssertEqual(t, Average(1, 2), 1.5)
	AssertEqual(t, Average[int](), 0.0)
	AssertEqual(t, Average(1.0, 2.0, 3.0), 2.0)
}

func TestPow(t *testing.T) {
	AssertEqual(t, Pow(2, 3), 8)
	AssertEqual(t, Pow(2, 0), 1)
	AssertEqual(t, Pow(5, 1), 5)
	AssertEqual(t, Pow(2, 10), 1024)
	AssertEqual(t, Pow(3, 4), 81)
	AssertEqual(t, Pow(2, -1), 0)
}

func TestIsPrime(t *testing.T) {
	AssertTrue(t, !IsPrime(0))
	AssertTrue(t, !IsPrime(1))
	AssertTrue(t, IsPrime(2))
	AssertTrue(t, IsPrime(3))
	AssertTrue(t, !IsPrime(4))
	AssertTrue(t, IsPrime(5))
	AssertTrue(t, !IsPrime(9))
	AssertTrue(t, IsPrime(11))
	AssertTrue(t, IsPrime(13))
	AssertTrue(t, !IsPrime(15))
	AssertTrue(t, IsPrime(17))
	AssertTrue(t, !IsPrime(100))
	AssertTrue(t, IsPrime(97))
}

func TestGCD(t *testing.T) {
	AssertEqual(t, GCD(12, 8), 4)
	AssertEqual(t, GCD(7, 3), 1)
	AssertEqual(t, GCD(0, 5), 5)
	AssertEqual(t, GCD(6, 6), 6)
	AssertEqual(t, GCD(54, 24), 6)
}

func TestLCM(t *testing.T) {
	AssertEqual(t, LCM(4, 6), 12)
	AssertEqual(t, LCM(7, 3), 21)
	AssertEqual(t, LCM(0, 5), 0)
	AssertEqual(t, LCM(6, 6), 6)
}

func TestFibonacci(t *testing.T) {
	AssertSliceEqual(t, Fibonacci(0), []int{})
	AssertSliceEqual(t, Fibonacci(1), []int{0})
	AssertSliceEqual(t, Fibonacci(2), []int{0, 1})
	AssertSliceEqual(t, Fibonacci(5), []int{0, 1, 1, 2, 3})
	AssertSliceEqual(t, Fibonacci(10), []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34})
}

func TestFibonacciNegative(t *testing.T) {
	AssertSliceEqual(t, Fibonacci(-1), []int{})
}

func TestRound(t *testing.T) {
	AssertEqual(t, Round(3.4), 3.0)
	AssertEqual(t, Round(3.5), 4.0)
	AssertEqual(t, Round(-3.4), -3.0)
	AssertEqual(t, Round(-3.5), -4.0)
	AssertEqual(t, Round(0.0), 0.0)
}

func TestFloor(t *testing.T) {
	AssertEqual(t, Floor(3.7), 3.0)
	AssertEqual(t, Floor(-3.7), -4.0)
	AssertEqual(t, Floor(-3.0), -3.0)
	AssertEqual(t, Floor(0.0), 0.0)
}

func TestCeil(t *testing.T) {
	AssertEqual(t, Ceil(3.2), 4.0)
	AssertEqual(t, Ceil(-3.2), -3.0)
	AssertEqual(t, Ceil(3.0), 3.0)
	AssertEqual(t, Ceil(0.0), 0.0)
}

func TestMinSlice(t *testing.T) {
	AssertEqual(t, MinSlice([]int{3, 1, 4, 1, 5}), 1)
	AssertEqual(t, MinSlice([]int{5, 4, 3, 2, 1}), 1)
	AssertEqual(t, MinSlice([]string{"b", "a", "c"}), "a")
	AssertEqual(t, MinSlice([]int{}), 0)
}

func TestMaxSlice(t *testing.T) {
	AssertEqual(t, MaxSlice([]int{3, 1, 4, 1, 5}), 5)
	AssertEqual(t, MaxSlice([]int{1, 2, 3, 4, 5}), 5)
	AssertEqual(t, MaxSlice([]string{"a", "b", "c"}), "c")
	AssertEqual(t, MaxSlice([]int{}), 0)
}

func TestFactorial(t *testing.T) {
	AssertEqual(t, Factorial(0), 1)
	AssertEqual(t, Factorial(1), 1)
	AssertEqual(t, Factorial(5), 120)
	AssertEqual(t, Factorial(10), 3628800)
	AssertEqual(t, Factorial(-1), 0)
}

func TestMedian(t *testing.T) {
	AssertEqual(t, Median(1, 2, 3, 4, 5), 3.0)
	AssertEqual(t, Median(1, 2, 3, 4), 2.5)
	AssertEqual(t, Median(1.5, 2.5, 3.5), 2.5)
	AssertEqual(t, Median[int](), 0.0)
}
