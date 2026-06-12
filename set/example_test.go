package set

import "fmt"

func ExampleNew() {
	s := New(1, 2, 3, 2, 1)
	fmt.Println(s.Len())
	fmt.Println(s.Contains(1))
	fmt.Println(s.Contains(4))
	// Output:
	// 3
	// true
	// false
}

func ExampleSet_Union() {
	s1 := New(1, 2, 3)
	s2 := New(3, 4, 5)
	u := s1.Union(s2)
	fmt.Println(u.Len())
	fmt.Println(u.Contains(1))
	fmt.Println(u.Contains(5))
	// Output:
	// 5
	// true
	// true
}

func ExampleSet_Intersection() {
	s1 := New(1, 2, 3, 4)
	s2 := New(3, 4, 5, 6)
	i := s1.Intersection(s2)
	fmt.Println(i.Len())
	fmt.Println(i.Contains(3))
	fmt.Println(i.Contains(1))
	// Output:
	// 2
	// true
	// false
}

func ExampleSet_Difference() {
	s1 := New(1, 2, 3, 4)
	s2 := New(3, 4, 5)
	d := s1.Difference(s2)
	fmt.Println(d.ToSlice())
	// Output:
	// [1 2]
}
