package json

import (
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestMarshal(t *testing.T) {
	s, err := Marshal(map[string]int{"a": 1})
	AssertTrue(t, err == nil)
	AssertEqual(t, string(s), `{"a":1}`)
}

func TestMustMarshal(t *testing.T) {
	AssertEqual(t, string(MustMarshal(map[string]int{"a": 1})), `{"a":1}`)
}

func TestUnmarshal(t *testing.T) {
	var m map[string]int
	err := Unmarshal(`{"a":1}`, &m)
	AssertTrue(t, err == nil)
	AssertEqual(t, m["a"], 1)
}

func TestMustUnmarshal(t *testing.T) {
	var m map[string]int
	MustUnmarshal(`{"a":1}`, &m)
	AssertEqual(t, m["a"], 1)
}

func TestMustMarshalPanic(t *testing.T) {
	AssertPanic(t, func() {
		MustMarshal(make(chan int))
	})
}

func TestMustUnmarshalPanic(t *testing.T) {
	AssertPanic(t, func() {
		var m map[string]int
		MustUnmarshal(`{invalid}`, &m)
	})
}

func TestMarshalIndent(t *testing.T) {
	s, _ := MarshalIndent(map[string]int{"a": 1}, "", "  ")
	AssertEqual(t, string(s), "{\n  \"a\": 1\n}")
}