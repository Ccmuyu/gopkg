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

func TestClone(t *testing.T) {
	m := map[string]int{"a": 1}
	cloned := Clone(m)
	AssertEqual(t, cloned["a"], 1)
	cloned["b"] = 2
	AssertTrue(t, m["b"] == 0)
}