package maps

import (
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	keys := Keys(m)
	AssertEqual(t, len(keys), 2)
}

func TestValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	values := Values(m)
	AssertEqual(t, len(values), 2)
}

func TestHasKey(t *testing.T) {
	m := map[string]int{"a": 1}
	AssertTrue(t, HasKey(m, "a"))
	AssertTrue(t, !HasKey(m, "b"))
}

func TestGet(t *testing.T) {
	m := map[string]int{"a": 1}
	AssertEqual(t, Get(m, "a"), 1)
}

func TestGetOk(t *testing.T) {
	m := map[string]int{"a": 1}
	v, ok := GetOk(m, "a")
	AssertTrue(t, ok)
	AssertEqual(t, v, 1)

	_, ok = GetOk(m, "b")
	AssertTrue(t, !ok)
}

func TestFilter(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	filtered := Filter(m, func(k string, v int) bool {
		return v > 1
	})
	AssertEqual(t, len(filtered), 2)
}

func TestMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	mapped := Map(m, func(k string, v int) int {
		return v * 2
	})
	AssertEqual(t, mapped["a"], 2)
	AssertEqual(t, mapped["b"], 4)
}

func TestMerge(t *testing.T) {
	m1 := map[string]int{"a": 1}
	m2 := map[string]int{"b": 2}
	merged := Merge(m1, m2)
	AssertEqual(t, len(merged), 2)
	AssertEqual(t, merged["a"], 1)
	AssertEqual(t, merged["b"], 2)
}

func TestMergeWith(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	merged := MergeWith(func(v1, v2 int) int {
		if v1 > v2 {
			return v1
		}
		return v2
	}, m1, m2)
	AssertEqual(t, len(merged), 3)
	AssertEqual(t, merged["a"], 1)
	AssertEqual(t, merged["b"], 3)
	AssertEqual(t, merged["c"], 4)
}

func TestMergeWithSum(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	merged := MergeWith(func(v1, v2 int) int {
		return v1 + v2
	}, m1, m2)
	AssertEqual(t, merged["a"], 1)
	AssertEqual(t, merged["b"], 5)
	AssertEqual(t, merged["c"], 4)
}

func TestClone(t *testing.T) {
	m := map[string]int{"a": 1}
	cloned := Clone(m)
	AssertEqual(t, cloned["a"], 1)
	cloned["b"] = 2
	AssertTrue(t, m["b"] == 0)
}

func TestPick(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	picked := Pick(m, "a", "c")
	AssertEqual(t, len(picked), 2)
	AssertEqual(t, picked["a"], 1)
	AssertEqual(t, picked["c"], 3)
	AssertTrue(t, !HasKey(picked, "b"))
}

func TestPickNonExistent(t *testing.T) {
	m := map[string]int{"a": 1}
	picked := Pick(m, "a", "b")
	AssertEqual(t, len(picked), 1)
	AssertEqual(t, picked["a"], 1)
}

func TestOmit(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	omitted := Omit(m, "a", "c")
	AssertEqual(t, len(omitted), 1)
	AssertEqual(t, omitted["b"], 2)
	AssertTrue(t, !HasKey(omitted, "a"))
}

func TestOmitAll(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	omitted := Omit(m, "a", "b")
	AssertEqual(t, len(omitted), 0)
}

func TestInvert(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	inverted := Invert(m)
	AssertEqual(t, len(inverted), 2)
	AssertEqual(t, inverted[1], "a")
	AssertEqual(t, inverted[2], "b")
}

func TestForEach(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	sum := 0
	ForEach(m, func(k string, v int) {
		sum += v
	})
	AssertEqual(t, sum, 6)
}

func TestIsEqual(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 1, "b": 2}
	AssertTrue(t, IsEqual(m1, m2))
}

func TestIsEqualNotEqual(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 1, "b": 3}
	AssertTrue(t, !IsEqual(m1, m2))
}

func TestIsEqualDifferentLen(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 1}
	AssertTrue(t, !IsEqual(m1, m2))
}

func TestMapKey(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	mapped := MapKey(m, func(k string) string {
		return "prefix_" + k
	})
	AssertEqual(t, len(mapped), 3)
	AssertEqual(t, mapped["prefix_a"], 1)
	AssertEqual(t, mapped["prefix_b"], 2)
	AssertEqual(t, mapped["prefix_c"], 3)
}