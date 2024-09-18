package sorts

import "testing"

func TestSort(t *testing.T) {

	var arr = []int{3, 4, 1, 4, 5, 2}

	Sort(arr, true)
	t.Log(arr)
	Sort(arr, false)
	t.Log(arr)
}
