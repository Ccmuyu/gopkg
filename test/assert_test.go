package test

import "testing"

func TestAssertEqualSupportsCompositeValues(t *testing.T) {
	AssertEqual(t, []int{1, 2}, []int{1, 2})
	AssertEqual(t, map[string]int{"a": 1}, map[string]int{"a": 1})
}

func TestAssertSliceEqualSupportsNestedSlices(t *testing.T) {
	AssertSliceEqual(t, [][]int{{1, 2}}, [][]int{{1, 2}})
}
