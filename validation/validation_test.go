package validation

import (
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestIsEmail(t *testing.T) {
	AssertTrue(t, IsEmail("test@example.com"))
	AssertTrue(t, !IsEmail("invalid"))
	AssertTrue(t, !IsEmail("test@"))
}

func TestIsPhone(t *testing.T) {
	AssertTrue(t, IsPhone("13812345678"))
	AssertTrue(t, !IsPhone("12345678901"))
	AssertTrue(t, !IsPhone("1234567890"))
}

func TestIsURL(t *testing.T) {
	AssertTrue(t, IsURL("https://example.com"))
	AssertTrue(t, IsURL("http://example.com"))
	AssertTrue(t, !IsURL("not a url"))
}

func TestIsIP(t *testing.T) {
	AssertTrue(t, IsIP("192.168.1.1"))
	AssertTrue(t, IsIP("::1"))
	AssertTrue(t, !IsIP("not an ip"))
}

func TestIsIPv4(t *testing.T) {
	AssertTrue(t, IsIPv4("192.168.1.1"))
	AssertTrue(t, !IsIPv4("192.168.1.256"))
	AssertTrue(t, !IsIPv4("::1"))
}

func TestIsAlpha(t *testing.T) {
	AssertTrue(t, IsAlpha("hello"))
	AssertTrue(t, !IsAlpha("hello1"))
}

func TestIsNumeric(t *testing.T) {
	AssertTrue(t, IsNumeric("12345"))
	AssertTrue(t, !IsNumeric("123a"))
}

func TestIsAlphanumeric(t *testing.T) {
	AssertTrue(t, IsAlphanumeric("hello123"))
	AssertTrue(t, !IsAlphanumeric("hello!"))
}

func TestIsMatch(t *testing.T) {
	AssertTrue(t, IsMatch("hello", "^hello$"))
	AssertTrue(t, !IsMatch("world", "^hello$"))
}

func TestIsEmailAddr(t *testing.T) {
	AssertTrue(t, IsEmailAddr("test@example.com"))
	AssertTrue(t, !IsEmailAddr("invalid"))
}

func TestIsPrivateIP(t *testing.T) {
	AssertTrue(t, IsPrivateIP("192.168.1.1"))
	AssertTrue(t, IsPrivateIP("10.0.0.1"))
	AssertTrue(t, IsPrivateIP("127.0.0.1"))
	AssertTrue(t, !IsPrivateIP("8.8.8.8"))
}