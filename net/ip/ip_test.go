package ip

import (
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestIsValid(t *testing.T) {
	AssertTrue(t, IsValid("192.168.1.1"))
	AssertTrue(t, IsValid("::1"))
	AssertTrue(t, !IsValid("invalid"))
}

func TestIsPrivate(t *testing.T) {
	AssertTrue(t, IsPrivate("192.168.1.1"))
	AssertTrue(t, IsPrivate("10.0.0.1"))
	AssertTrue(t, IsPrivate("127.0.0.1"))
	AssertTrue(t, !IsPrivate("8.8.8.8"))
}

func TestToInt(t *testing.T) {
	n, err := ToInt("192.168.1.1")
	AssertTrue(t, err == nil)
	AssertEqual(t, n, int64(3232235777))
}

func TestToIntInvalid(t *testing.T) {
	_, err := ToInt("invalid")
	AssertTrue(t, err != nil)
}

func TestToIntIPv6(t *testing.T) {
	_, err := ToInt("::1")
	AssertTrue(t, err != nil)
}

func TestIntToIP(t *testing.T) {
	AssertEqual(t, IntToIP(3232235777), "192.168.1.1")
}

func TestToJSON(t *testing.T) {
	data, _ := ToJSON("192.168.1.1")
	AssertTrue(t, len(data) > 0)
}