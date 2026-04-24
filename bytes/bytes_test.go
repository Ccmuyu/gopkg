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

func TestHumanReadable(t *testing.T) {
	AssertEqual(t, HumanReadable(500), "500 B")
	AssertEqual(t, HumanReadable(1024), "1 KB")
	AssertEqual(t, HumanReadable(1536), "1.50 KB")
	AssertEqual(t, HumanReadable(1048576), "1 MB")
}