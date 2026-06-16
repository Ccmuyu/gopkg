package bytes

import (
	"encoding/hex"
	"encoding/base64"
	"math"
	"strconv"
	"strings"
)

func HexEncode(src []byte) string {
	return hex.EncodeToString(src)
}

func HexDecode(s string) ([]byte, error) {
	return hex.DecodeString(s)
}

func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func Base64Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func Base64URLEncode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func Base64URLDecode(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(s)
}

func Base64EncodeString(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64DecodeString(s string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Base64URLEncodeString(s string) string {
	return base64.URLEncoding.EncodeToString([]byte(s))
}

func Base64URLDecodeString(s string) (string, error) {
	data, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func HumanReadable(bytes int64) string {
	const unit = 1024
	if bytes < 0 {
		if bytes == math.MinInt64 {
			return "-9223372036854775808 B"
		}
		return "-" + HumanReadable(-bytes)
	}
	if bytes < unit {
		return intToStr(bytes) + " B"
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return floatToStr(float64(bytes)/float64(div)) + " " + units[exp] + "B"
}

var units = []string{"K", "M", "G", "T", "P", "E"}

func intToStr(n int64) string {
	if n == math.MinInt64 {
		return "-9223372036854775808"
	}
	if n == 0 {
		return "0"
	}
	var result []byte
	negative := n < 0
	if negative {
		n = -n
	}
	for n > 0 {
		result = append([]byte{byte('0' + n%10)}, result...)
		n /= 10
	}
	if negative {
		result = append([]byte{'-'}, result...)
	}
	return string(result)
}

func floatToStr(f float64) string {
	s := strconv.FormatFloat(f, 'f', 2, 64)
	s = strings.TrimRight(s, "0")
	s = strings.TrimRight(s, ".")
	return s
}