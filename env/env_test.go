package env

import (
	"os"
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func setenvCleanup(t *testing.T, key, value string) {
	prev, existed := os.LookupEnv(key)
	os.Setenv(key, value)
	t.Cleanup(func() {
		if existed {
			os.Setenv(key, prev)
		} else {
			os.Unsetenv(key)
		}
	})
}

func TestGet(t *testing.T) {
	setenvCleanup(t, "TEST_KEY", "hello")
	AssertEqual(t, Get("TEST_KEY", "default"), "hello")
	AssertEqual(t, Get("NONEXISTENT_KEY_XYZ", "default"), "default")
}

func TestMustGet(t *testing.T) {
	setenvCleanup(t, "TEST_KEY2", "world")
	AssertEqual(t, MustGet("TEST_KEY2"), "world")

	defer func() {
		AssertTrue(t, recover() != nil)
	}()
	MustGet("NONEXISTENT_KEY_MUST")
}

func TestSet(t *testing.T) {
	err := Set("TEST_KEY3", "value3")
	AssertTrue(t, err == nil)
	AssertEqual(t, os.Getenv("TEST_KEY3"), "value3")
	os.Unsetenv("TEST_KEY3")
}

func TestUnset(t *testing.T) {
	setenvCleanup(t, "TEST_KEY4", "value4")
	Unset("TEST_KEY4")
	AssertEqual(t, os.Getenv("TEST_KEY4"), "")
}

func TestHas(t *testing.T) {
	setenvCleanup(t, "TEST_KEY5", "value5")
	AssertTrue(t, Has("TEST_KEY5"))
	AssertTrue(t, !Has("NONEXISTENT_KEY_HAS"))
}

func TestEnvironMap(t *testing.T) {
	setenvCleanup(t, "TEST_KEY6", "value6")
	m := EnvironMap()
	AssertEqual(t, m["TEST_KEY6"], "value6")
	AssertTrue(t, len(m) > 0)
}
