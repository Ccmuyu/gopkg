package sorts

import "sort"

type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Sort[T Comparable](arr []T, asc bool) {
	sort.Slice(arr, func(i, j int) bool {
		if asc {
			return arr[i] < arr[j]
		}
		return arr[i] > arr[j]
	})

}
