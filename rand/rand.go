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
	// 使用无符号差值避免 max-min 在跨越整个有符号整数范围时溢出。
	span := uint64(max) - uint64(min)
	return int(uint64(min) + uint64n(span))
}

func uint64n(n uint64) uint64 {
	// 拒绝采样消除取模偏差。Int 保证 n > 0。
	threshold := -n % n
	for {
		v := rand.Uint64()
		if v >= threshold {
			return v % n
		}
	}
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
