package slices

import (
	"testing"

	gtest "github.com/Ccmuyu/gopkg/test"
)

func TestDedup(t *testing.T) {
	gtest.AssertSliceEqual(t, Dedup([]int{1, 2, 2, 3, 3, 3}), []int{1, 2, 3})
	gtest.AssertSliceEqual(t, Dedup([]string{"a", "b", "a"}), []string{"a", "b"})
	gtest.AssertSliceEqual(t, Dedup([]int{}), []int{})
}

func TestFilter(t *testing.T) {
	res := Filter([]int{1, 2, 3, 4, 5}, func(v int) bool {
		return v%2 == 0
	})
	gtest.AssertSliceEqual(t, res, []int{2, 4})
}

func TestMap(t *testing.T) {
	res := Map([]int{1, 2, 3}, func(v int) int {
		return v * 2
	})
	gtest.AssertSliceEqual(t, res, []int{2, 4, 6})
}

func TestFlatMap(t *testing.T) {
	res := FlatMap([]int{1, 2, 3}, func(v int) []int {
		return []int{v, v * 2}
	})
	gtest.AssertSliceEqual(t, res, []int{1, 2, 2, 4, 3, 6})
}

func TestMerge(t *testing.T) {
	res := Merge([]int{1, 2}, []int{3, 4}, []int{5})
	gtest.AssertSliceEqual(t, res, []int{1, 2, 3, 4, 5})
}

func TestContains(t *testing.T) {
	gtest.AssertTrue(t, Contains([]int{1, 2, 3}, 2))
	gtest.AssertTrue(t, !Contains([]int{1, 2, 3}, 4))
}

func TestIndex(t *testing.T) {
	gtest.AssertEqual(t, Index([]int{1, 2, 3}, 2), 1)
	gtest.AssertEqual(t, Index([]int{1, 2, 3}, 4), -1)
}

func TestSplit(t *testing.T) {
	res := Split([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	gtest.AssertEqual(t, len(res), 3)
	gtest.AssertEqual(t, len(res[2]), 1)

	res = Split([]int{1, 2, 3, 4, 5, 6}, 3)
	gtest.AssertEqual(t, len(res), 2)
}

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	Reverse(arr)
	gtest.AssertSliceEqual(t, arr, []int{5, 4, 3, 2, 1})
}

func TestGroupBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 25},
		{"Diana", 30},
	}
	result := GroupBy(people, func(p Person) int {
		return p.Age
	})
	gtest.AssertEqual(t, len(result), 2)
	gtest.AssertEqual(t, len(result[25]), 2)
	gtest.AssertEqual(t, len(result[30]), 2)
}