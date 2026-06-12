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

func TestReduce(t *testing.T) {
	sum := Reduce([]int{1, 2, 3, 4, 5}, 0, func(acc, v int) int {
		return acc + v
	})
	gtest.AssertEqual(t, sum, 15)
}

func TestSome(t *testing.T) {
	gtest.AssertTrue(t, Some([]int{1, 2, 3, 4, 5}, func(v int) bool {
		return v > 3
	}))
	gtest.AssertTrue(t, !Some([]int{1, 2, 3}, func(v int) bool {
		return v > 5
	}))
	gtest.AssertTrue(t, !Some([]int{}, func(v int) bool {
		return true
	}))
}

func TestEvery(t *testing.T) {
	gtest.AssertTrue(t, Every([]int{2, 4, 6, 8}, func(v int) bool {
		return v%2 == 0
	}))
	gtest.AssertTrue(t, !Every([]int{2, 4, 5, 8}, func(v int) bool {
		return v%2 == 0
	}))
	gtest.AssertTrue(t, Every([]int{}, func(v int) bool {
		return false
	}))
}

func TestNone(t *testing.T) {
	gtest.AssertTrue(t, None([]int{1, 3, 5}, func(v int) bool {
		return v%2 == 0
	}))
	gtest.AssertTrue(t, !None([]int{1, 2, 3}, func(v int) bool {
		return v%2 == 0
	}))
}

func TestWithout(t *testing.T) {
	gtest.AssertSliceEqual(t, Without([]int{1, 2, 3, 4, 5}, 2, 4), []int{1, 3, 5})
	gtest.AssertSliceEqual(t, Without([]int{1, 1, 2, 2}, 1), []int{2, 2})
	gtest.AssertSliceEqual(t, Without([]int{}, 1), []int{})
}

func TestIntersection(t *testing.T) {
	gtest.AssertSliceEqual(t, Intersection([]int{1, 2, 3, 4}, []int{3, 4, 5, 6}), []int{3, 4})
	gtest.AssertEqual(t, len(Intersection([]int{1, 2}, []int{3, 4})), 0)
}

func TestUnion(t *testing.T) {
	gtest.AssertSliceEqual(t, Union([]int{1, 2, 3}, []int{3, 4, 5}), []int{1, 2, 3, 4, 5})
	gtest.AssertSliceEqual(t, Union([]int{1, 2}, []int{2, 3}, []int{3, 4}), []int{1, 2, 3, 4})
}

func TestDifference(t *testing.T) {
	a, b := Difference([]int{1, 2, 3, 4}, []int{3, 4, 5, 6})
	gtest.AssertSliceEqual(t, a, []int{1, 2})
	gtest.AssertSliceEqual(t, b, []int{5, 6})
}

func TestChunk(t *testing.T) {
	chunks := Chunk([]int{1, 2, 3, 4, 5, 6}, 3)
	gtest.AssertEqual(t, len(chunks), 2)
	gtest.AssertSliceEqual(t, chunks[0], []int{1, 2, 3})
	gtest.AssertSliceEqual(t, chunks[1], []int{4, 5, 6})

	chunks2 := Chunk([]int{1, 2, 3, 4, 5}, 2)
	gtest.AssertEqual(t, len(chunks2), 3)
}

func TestSort(t *testing.T) {
	arr := []int{3, 1, 4, 1, 5}
	Sort(arr)
	gtest.AssertSliceEqual(t, arr, []int{1, 1, 3, 4, 5})
}

func TestFill(t *testing.T) {
	arr := make([]int, 5)
	Fill(arr, 42)
	gtest.AssertSliceEqual(t, arr, []int{42, 42, 42, 42, 42})
}

func TestTake(t *testing.T) {
	gtest.AssertSliceEqual(t, Take([]int{1, 2, 3, 4, 5}, 3), []int{1, 2, 3})
	gtest.AssertSliceEqual(t, Take([]int{1, 2, 3, 4, 5}, 10), []int{1, 2, 3, 4, 5})
	gtest.AssertSliceEqual(t, Take([]int{1, 2, 3}, 0), []int{})
	gtest.AssertSliceEqual(t, Take([]int{1, 2, 3}, -1), []int{})
}

func TestDrop(t *testing.T) {
	gtest.AssertSliceEqual(t, Drop([]int{1, 2, 3, 4, 5}, 2), []int{3, 4, 5})
	gtest.AssertSliceEqual(t, Drop([]int{1, 2, 3}, 10), []int{})
	gtest.AssertSliceEqual(t, Drop([]int{1, 2, 3}, 0), []int{1, 2, 3})
}