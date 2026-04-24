package slices

import (
	"reflect"
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func AssertSliceEqual[T any](t *testing.T, actual, expected []T) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual: %v\nexpected: %v", actual, expected)
	}
}

func TestDedup(t *testing.T) {
	AssertSliceEqual(t, Dedup([]int{1, 2, 2, 3, 3, 3}), []int{1, 2, 3})
	AssertSliceEqual(t, Dedup([]string{"a", "b", "a"}), []string{"a", "b"})
	AssertSliceEqual(t, Dedup([]int{}), []int{})
}

func TestFilter(t *testing.T) {
	res := Filter([]int{1, 2, 3, 4, 5}, func(v int) bool {
		return v%2 == 0
	})
	AssertSliceEqual(t, res, []int{2, 4})
}

func TestMap(t *testing.T) {
	res := Map([]int{1, 2, 3}, func(v int) int {
		return v * 2
	})
	AssertSliceEqual(t, res, []int{2, 4, 6})
}

func TestFlatMap(t *testing.T) {
	res := FlatMap([]int{1, 2, 3}, func(v int) []int {
		return []int{v, v * 2}
	})
	AssertSliceEqual(t, res, []int{1, 2, 2, 4, 3, 6})
}

func TestMerge(t *testing.T) {
	res := Merge([]int{1, 2}, []int{3, 4}, []int{5})
	AssertSliceEqual(t, res, []int{1, 2, 3, 4, 5})
}

func TestContains(t *testing.T) {
	AssertTrue(t, Contains([]int{1, 2, 3}, 2))
	AssertTrue(t, !Contains([]int{1, 2, 3}, 4))
}

func TestIndex(t *testing.T) {
	AssertEqual(t, Index([]int{1, 2, 3}, 2), 1)
	AssertEqual(t, Index([]int{1, 2, 3}, 4), -1)
}