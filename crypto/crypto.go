package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/base64"
	"hash"
	"hash/crc32"
)

func MD5(s string) string {
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}

func MD5Bytes(data []byte) string {
	h := md5.Sum(data)
	return hex.EncodeToString(h[:])
}

func SHA1(s string) string {
	h := sha1.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}

func SHA1Bytes(data []byte) string {
	h := sha1.Sum(data)
	return hex.EncodeToString(h[:])
}

func SHA256(s string) string {
	h := sha256.Sum256([]byte(s))
	return hex.EncodeToString(h[:])
}

func SHA256Bytes(data []byte) string {
	h := sha256.Sum256(data)
	return hex.EncodeToString(h[:])
}

func SHA512(s string) string {
	h := sha512.Sum512([]byte(s))
	return hex.EncodeToString(h[:])
}

func SHA512Bytes(data []byte) string {
	h := sha512.Sum512(data)
	return hex.EncodeToString(h[:])
}

func HMAC(algorithm func() hash.Hash, key, data []byte) string {
	mac := hmac.New(algorithm, key)
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil))
}

func CRC32(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

func CRC32String(s string) uint32 {
	return crc32.ChecksumIEEE([]byte(s))
}

func RandomToken(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
