package slices

import "fmt"

func ExampleReduce() {
	sum := Reduce([]int{1, 2, 3, 4, 5}, 0, func(acc, v int) int {
		return acc + v
	})
	fmt.Println(sum)
	// Output:
	// 15
}

func ExampleSome() {
	result := Some([]int{1, 2, 3, 4, 5}, func(v int) bool {
		return v > 3
	})
	fmt.Println(result)
	// Output:
	// true
}

func ExampleEvery() {
	result := Every([]int{2, 4, 6, 8}, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(result)
	// Output:
	// true
}

func ExampleUnion() {
	result := Union([]int{1, 2, 3}, []int{3, 4, 5})
	fmt.Println(result)
	// Output:
	// [1 2 3 4 5]
}
