package fn

import (
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestAnd(t *testing.T) {
	isPositive := func(v int) bool { return v > 0 }
	isEven := func(v int) bool { return v%2 == 0 }
	positiveAndEven := And(isPositive, isEven)

	AssertTrue(t, positiveAndEven(2))
	AssertTrue(t, !positiveAndEven(1))
	AssertTrue(t, !positiveAndEven(-2))
}

func TestOr(t *testing.T) {
	isNegative := func(v int) bool { return v < 0 }
	isZero := func(v int) bool { return v == 0 }
	nonPositive := Or(isNegative, isZero)

	AssertTrue(t, nonPositive(-1))
	AssertTrue(t, nonPositive(0))
	AssertTrue(t, !nonPositive(1))
}

func TestNot(t *testing.T) {
	isZero := func(v int) bool { return v == 0 }
	notZero := Not(isZero)

	AssertTrue(t, notZero(1))
	AssertTrue(t, !notZero(0))
}

func TestAndEmpty(t *testing.T) {
	always := And[int]()
	AssertTrue(t, always(42))
}

func TestOrEmpty(t *testing.T) {
	never := Or[int]()
	AssertTrue(t, !never(42))
}

func TestTuple2(t *testing.T) {
	tup := MakeTuple2("hello", 42)
	AssertEqual(t, tup.A, "hello")
	AssertEqual(t, tup.B, 42)
}

func TestTuple3(t *testing.T) {
	tup := MakeTuple3(1, "two", 3.0)
	AssertEqual(t, tup.A, 1)
	AssertEqual(t, tup.B, "two")
	AssertEqual(t, tup.C, 3.0)
}
