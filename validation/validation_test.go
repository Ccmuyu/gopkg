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

func TestIsUUID(t *testing.T) {
	AssertTrue(t, IsUUID("550e8400-e29b-41d4-a716-446655440000"))
	AssertTrue(t, !IsUUID("not-a-uuid"))
	AssertTrue(t, !IsUUID("550e8400-e29b-41d4-a716-44665544000z"))
}

func TestIsMAC(t *testing.T) {
	AssertTrue(t, IsMAC("00:1A:2B:3C:4D:5E"))
	AssertTrue(t, IsMAC("00-1A-2B-3C-4D-5E"))
	AssertTrue(t, !IsMAC("not-a-mac"))
}

func TestIsJSON(t *testing.T) {
	AssertTrue(t, IsJSON(`{"key":"value"}`))
	AssertTrue(t, IsJSON(`[1,2,3]`))
	AssertTrue(t, IsJSON(`"string"`))
	AssertTrue(t, !IsJSON(`{invalid}`))
}

func TestIsHexColor(t *testing.T) {
	AssertTrue(t, IsHexColor("#fff"))
	AssertTrue(t, IsHexColor("#FFFFFF"))
	AssertTrue(t, IsHexColor("#000000"))
	AssertTrue(t, !IsHexColor("not-a-color"))
	AssertTrue(t, !IsHexColor("#GGG"))
}

func TestIsPostalCode(t *testing.T) {
	AssertTrue(t, IsPostalCode("100000"))
	AssertTrue(t, IsPostalCode("200000"))
	AssertTrue(t, !IsPostalCode("12345"))
	AssertTrue(t, !IsPostalCode("abcdef"))
}

func TestIsPort(t *testing.T) {
	AssertTrue(t, IsPort(80))
	AssertTrue(t, IsPort(443))
	AssertTrue(t, IsPort(1))
	AssertTrue(t, IsPort(65535))
	AssertTrue(t, !IsPort(0))
	AssertTrue(t, !IsPort(-1))
	AssertTrue(t, !IsPort(65536))
}

func TestIsLatitude(t *testing.T) {
	AssertTrue(t, IsLatitude(0))
	AssertTrue(t, IsLatitude(90))
	AssertTrue(t, IsLatitude(-90))
	AssertTrue(t, !IsLatitude(91))
	AssertTrue(t, !IsLatitude(-91))
}

func TestIsLongitude(t *testing.T) {
	AssertTrue(t, IsLongitude(0))
	AssertTrue(t, IsLongitude(180))
	AssertTrue(t, IsLongitude(-180))
	AssertTrue(t, !IsLongitude(181))
	AssertTrue(t, !IsLongitude(-181))
}

func TestIsCreditCard(t *testing.T) {
	AssertTrue(t, IsCreditCard("4111111111111111"))
	AssertTrue(t, IsCreditCard("5500000000000004"))
	AssertTrue(t, !IsCreditCard("1234567890123456"))
	AssertTrue(t, !IsCreditCard("1234"))
	AssertTrue(t, !IsCreditCard("abcdefghijklmnop"))
}