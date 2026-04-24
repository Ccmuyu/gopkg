package str

import (
	"strconv"
)

func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func Atoi64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}

func Atoi32(s string) int32 {
	v, _ := strconv.ParseInt(s, 10, 32)
	return int32(v)
}

func Itoa(i int) string {
	return strconv.Itoa(i)
}

func Itoa64(i int64) string {
	return strconv.FormatInt(i, 10)
}

func ParseFloat(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

func FormatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func FormatFloatPrec(f float64, prec int) string {
	return strconv.FormatFloat(f, 'f', prec, 64)
}

func ParseBool(s string) bool {
	v, _ := strconv.ParseBool(s)
	return v
}

func FormatBool(b bool) string {
	return strconv.FormatBool(b)
}