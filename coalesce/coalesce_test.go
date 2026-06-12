package coalesce

import (
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestCoalesceInt(t *testing.T) {
	AssertEqual(t, Coalesce(0, 0, 5, 3), 5)
	AssertEqual(t, Coalesce(0, 0, 0), 0)
	AssertEqual(t, Coalesce(1, 2, 3), 1)
}

func TestCoalesceString(t *testing.T) {
	AssertEqual(t, Coalesce("", "", "hello", "world"), "hello")
	AssertEqual(t, Coalesce("", "", ""), "")
	AssertEqual(t, Coalesce("first", "second"), "first")
}

func TestCoalesceEmpty(t *testing.T) {
	AssertEqual(t, Coalesce[int](), 0)
	AssertEqual(t, Coalesce[string](), "")
}

func TestCoalesceSliceInt(t *testing.T) {
	AssertEqual(t, CoalesceSlice([]int{0, 0, 5, 3}), 5)
	AssertEqual(t, CoalesceSlice([]int{0, 0, 0}), 0)
}

func TestCoalesceSliceString(t *testing.T) {
	AssertEqual(t, CoalesceSlice([]string{"", "", "hello"}), "hello")
	AssertEqual(t, CoalesceSlice([]string{}), "")
}

func TestCoalescePtr(t *testing.T) {
	var nilPtr *int
	a := 42
	AssertEqual(t, Coalesce(nilPtr, nilPtr, &a), &a)
	AssertEqual(t, Coalesce[*int](nilPtr, nilPtr, nilPtr), nilPtr)
}
