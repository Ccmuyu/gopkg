package crypto

import (
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestMD5(t *testing.T) {
	AssertEqual(t, MD5("hello"), "5d41402abc4b2a76b9719d911017c592")
}

func TestMD5Bytes(t *testing.T) {
	AssertEqual(t, MD5Bytes([]byte("hello")), "5d41402abc4b2a76b9719d911017c592")
}

func TestSHA1(t *testing.T) {
	AssertEqual(t, SHA1("hello"), "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d")
}

func TestSHA1Bytes(t *testing.T) {
	AssertEqual(t, SHA1Bytes([]byte("hello")), "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d")
}

func TestSHA256(t *testing.T) {
	AssertEqual(t, SHA256("hello"), "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824")
}

func TestSHA256Bytes(t *testing.T) {
	AssertEqual(t, SHA256Bytes([]byte("hello")), "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824")
}