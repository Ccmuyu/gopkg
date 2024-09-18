package test

import "testing"

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
