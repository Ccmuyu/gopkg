package test

import (
	"reflect"
	"regexp"
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

func AssertNil(t *testing.T, v interface{}) {
	if v != nil && !reflect.ValueOf(v).IsNil() {
		t.Errorf("[AssertNil] expected nil, got %v", v)
	}
}

func AssertNotNil(t *testing.T, v interface{}) {
	if v == nil || reflect.ValueOf(v).IsNil() {
		t.Errorf("[AssertNotNil] expected non-nil, got nil")
	}
}

func AssertError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("[AssertError] expected error, got nil")
	}
}

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("[AssertNoError] expected no error, got %v", err)
	}
}

func AssertPanic(t *testing.T, fn func()) {
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		fn()
	}()
	if !panicked {
		t.Errorf("[AssertPanic] expected panic, but none occurred")
	}
}

func AssertMatch(t *testing.T, s, pattern string) {
	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		t.Errorf("[AssertMatch] invalid pattern %q: %v", pattern, err)
		return
	}
	if !matched {
		t.Errorf("[AssertMatch] %q does not match pattern %q", s, pattern)
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

func SliceContains[T comparable](arr []T, target T) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}
