package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
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

// RandomToken 返回由 n 个加密安全随机字节编码的 URL-safe Base64 字符串。
// n <= 0 或系统随机源失败时返回空字符串。
func RandomToken(n int) string {
	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
