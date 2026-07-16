package rand

import (
	crand "crypto/rand"
	"math/rand/v2"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// String 返回长度为 n 的随机字母数字字符串。
// 注意：基于 math/rand，非加密安全，切勿用于生成密钥/令牌/密码；
// 此类场景请使用 crypto.RandomToken。
func String(n int) string {
	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[Int(0, len(letterBytes))]
	}
	return string(b)
}

func Int(min, max int) int {
	if min >= max {
		return min
	}
	return rand.IntN(max-min) + min
}

func Bytes(n int) []byte {
	if n <= 0 {
		return nil
	}
	b := make([]byte, n)
	if _, err := crand.Read(b); err != nil {
		return nil
	}
	return b
}

func Choice[T any](slice []T) T {
	var zero T
	if len(slice) == 0 {
		return zero
	}
	return slice[rand.IntN(len(slice))]
}

func Shuffle[T any](slice []T) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}
