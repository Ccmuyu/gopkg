package sorts

import (
	"testing"

	gtest "github.com/Ccmuyu/gopkg/test"
)

func TestSort(t *testing.T) {
	arr := []int{3, 4, 1, 4, 5, 2}
	Sort(arr, true)
	gtest.AssertSliceEqual(t, arr, []int{1, 2, 3, 4, 4, 5})
	Sort(arr, false)
	gtest.AssertSliceEqual(t, arr, []int{5, 4, 4, 3, 2, 1})
}

func TestSortFunc(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Bob", 30},
		{"Alice", 25},
		{"Charlie", 30},
	}
	SortFunc(people, func(a, b Person) int {
		if a.Age != b.Age {
			return a.Age - b.Age
		}
		if a.Name < b.Name {
			return -1
		}
		return 1
	})
	gtest.AssertEqual(t, people[0].Name, "Alice")
	gtest.AssertEqual(t, people[1].Name, "Bob")
	gtest.AssertEqual(t, people[2].Name, "Charlie")
}
