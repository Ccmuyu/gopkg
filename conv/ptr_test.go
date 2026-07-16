package conv

import (
	"testing"

	gtest "github.com/Ccmuyu/gopkg/test"
)

func TestVal(t *testing.T) {
	gtest.AssertEqual(t, Val[int](nil), 0)
	gtest.AssertEqual(t, Val[string](nil), "")
	gtest.AssertEqual(t, Val(Ptr(42)), 42)
	gtest.AssertEqual(t, Val(Ptr("hi")), "hi")
	gtest.AssertEqual(t, Val(Ptr(true)), true)
}

func TestPtr(t *testing.T) {
	p := Ptr(3.14)
	gtest.AssertTrue(t, p != nil)
	gtest.AssertEqual(t, *p, 3.14)
}

func TestValPtrRoundTrip(t *testing.T) {
	type point struct{ X, Y int }
	pt := point{X: 1, Y: 2}
	gtest.AssertEqual(t, Val(Ptr(pt)), pt)
}
