package test

import (
	"testing"
)

func AssertEqual(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Errorf("[AssertEqual] actual: %v\nexpected: %v", actual, expected)
		return
	}
}

func AssertNotEqual(t *testing.T, actual, expected interface{}) {
	if actual == expected {
		t.Errorf("[AssertNotEqual] actual: %v\nexpected: %v", actual, expected)
	}
}

func AssertTrue(t *testing.T, actual bool) {
	if !actual {
		t.Errorf("[AssertTrue] actual: %v\nexpected: %v", actual, true)
	}
}

func AssertSliceEqual[T any](t *testing.T, actual, expected []T) {
	if len(actual) != len(expected) {
		t.Errorf("SliceEqual: len mismatch: %d vs %d", len(actual), len(expected))
		return
	}
	for i := range actual {
		if any(actual[i]) != any(expected[i]) {
			t.Errorf("SliceEqual: index %d mismatch: %v vs %v", i, actual[i], expected[i])
		}
	}
}

func ContainsStr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
