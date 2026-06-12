package times

import "fmt"

func ExampleFormat() {
	t := Unix(1726675518, 0)
	fmt.Println(Format(t))
	// Output:
	// 2024-09-19 00:05:18
}

func ExampleBeginOfDay() {
	t := Unix(1726675518, 0)
	start := BeginOfDay(t)
	fmt.Println(Format(start))
	// Output:
	// 2024-09-19 00:00:00
}

func ExampleEndOfDay() {
	t := Unix(1726675518, 0)
	end := EndOfDay(t)
	fmt.Println(Format(end))
	// Output:
	// 2024-09-19 23:59:59
}

func ExampleIsWeekend() {
	saturday := Unix(1726876800, 0)
	fmt.Println(IsWeekend(saturday))
	// Output:
	// true
}

func ExampleDaysBetween() {
	start := Unix(1726675518, 0)
	end := Unix(1727280318, 0)
	fmt.Println(DaysBetween(start, end))
	// Output:
	// 7
}
