package slices

import (
	. "github.com/Ccmuyu/gopkg/test"
	"testing"
)

func TestSplit(t *testing.T) {

	res := Split([]int{1, 2, 3, 4, 5, 6, 7}, 3)

	AssertEqual(t, len(res), 3)
	AssertEqual(t, len(res[2]), 1)
	//t.Log(res)

	res = Split([]int{1, 2, 3, 4, 5, 6}, 3)
	AssertEqual(t, len(res), 2)
	//t.Log(res)
}
