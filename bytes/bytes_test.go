package bytes

import (
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestHexEncodeDecode(t *testing.T) {
	encoded := HexEncode([]byte("hello"))
	AssertEqual(t, encoded, "68656c6c6f")

	decoded, err := HexDecode(encoded)
	AssertTrue(t, err == nil)
	AssertEqual(t, string(decoded), "hello")
}

func TestBase64EncodeDecode(t *testing.T) {
	encoded := Base64Encode([]byte("hello"))
	AssertEqual(t, encoded, "aGVsbG8=")

	decoded, err := Base64Decode(encoded)
	AssertTrue(t, err == nil)
	AssertEqual(t, string(decoded), "hello")
}

func TestBase64URLEncodeDecode(t *testing.T) {
	encoded := Base64URLEncode([]byte("hello?world=foo"))
	AssertTrue(t, encoded != "")
	decoded, err := Base64URLDecode(encoded)
	AssertTrue(t, err == nil)
	AssertEqual(t, string(decoded), "hello?world=foo")
}

func TestBase64EncodeStringDecodeString(t *testing.T) {
	encoded := Base64EncodeString("hello")
	AssertEqual(t, encoded, "aGVsbG8=")

	decoded, err := Base64DecodeString(encoded)
	AssertTrue(t, err == nil)
	AssertEqual(t, decoded, "hello")
}

func TestBase64DecodeStringError(t *testing.T) {
	_, err := Base64DecodeString("invalid!!!")
	AssertTrue(t, err != nil)
}

func TestBase64URLEncodeStringDecodeString(t *testing.T) {
	encoded := Base64URLEncodeString("hello?world=")
	AssertEqual(t, encoded, "aGVsbG8_d29ybGQ9")

	decoded, err := Base64URLDecodeString(encoded)
	AssertTrue(t, err == nil)
	AssertEqual(t, decoded, "hello?world=")
}

func TestBase64URLDecodeStringError(t *testing.T) {
	_, err := Base64URLDecodeString("invalid!!!")
	AssertTrue(t, err != nil)
}

func TestHumanReadable(t *testing.T) {
	AssertEqual(t, HumanReadable(500), "500 B")
	AssertEqual(t, HumanReadable(1024), "1 KB")
	AssertEqual(t, HumanReadable(1536), "1.50 KB")
	AssertEqual(t, HumanReadable(1048576), "1 MB")
}