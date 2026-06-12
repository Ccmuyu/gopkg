package rand

import (
	crand "crypto/rand"
	"math/rand/v2"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

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
	_, _ = crand.Read(b)
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
