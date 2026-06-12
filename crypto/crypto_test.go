package crypto

import (
	"crypto/sha256"
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

func TestSHA512(t *testing.T) {
	AssertEqual(t, SHA512("hello"), "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043")
}

func TestSHA512Bytes(t *testing.T) {
	AssertEqual(t, SHA512Bytes([]byte("hello")), "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043")
}

func TestCRC32(t *testing.T) {
	AssertEqual(t, CRC32([]byte("hello")), uint32(907060870))
}

func TestCRC32String(t *testing.T) {
	AssertEqual(t, CRC32String("hello"), uint32(907060870))
}

func TestRandomToken(t *testing.T) {
	token := RandomToken(16)
	AssertTrue(t, len(token) > 0)
	token2 := RandomToken(16)
	AssertTrue(t, token != token2)
}

func TestHMAC(t *testing.T) {
	result := HMAC(sha256.New, []byte("key"), []byte("hello"))
	AssertEqual(t, len(result), 64)
	result2 := HMAC(sha256.New, []byte("key"), []byte("hello"))
	AssertEqual(t, result, result2)
}