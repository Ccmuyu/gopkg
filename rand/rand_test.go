package rand

import (
	"math"
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestString(t *testing.T) {
	s := String(10)
	AssertEqual(t, len(s), 10)
	AssertTrue(t, s != "")
}

func TestStringZero(t *testing.T) {
	AssertEqual(t, String(0), "")
	AssertEqual(t, String(-1), "")
}

func TestInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := Int(5, 10)
		AssertTrue(t, n >= 5)
		AssertTrue(t, n < 10)
	}
}

func TestIntMinMax(t *testing.T) {
	AssertEqual(t, Int(5, 5), 5)
	AssertEqual(t, Int(10, 5), 10)
}

func TestIntFullSignedRange(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := Int(math.MinInt, math.MaxInt)
		AssertTrue(t, n >= math.MinInt)
		AssertTrue(t, n < math.MaxInt)
	}
}

func TestBytes(t *testing.T) {
	b := Bytes(16)
	AssertEqual(t, len(b), 16)
}

func TestBytesZero(t *testing.T) {
	AssertEqual(t, len(Bytes(0)), 0)
	AssertEqual(t, len(Bytes(-1)), 0)
}

func TestChoice(t *testing.T) {
	slice := []int{10, 20, 30}
	seen := make(map[int]bool)
	for i := 0; i < 100; i++ {
		seen[Choice(slice)] = true
	}
	AssertTrue(t, seen[10])
	AssertTrue(t, seen[20])
	AssertTrue(t, seen[30])
}

func TestChoiceEmpty(t *testing.T) {
	AssertEqual(t, Choice([]int{}), 0)
	AssertEqual(t, Choice([]string{}), "")
}

func TestShuffle(t *testing.T) {
	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cp := make([]int, len(original))
	copy(cp, original)
	Shuffle(cp)
	AssertEqual(t, len(cp), 10)
	// must contain same elements
	sum := 0
	for _, v := range cp {
		sum += v
	}
	AssertEqual(t, sum, 55)
}
