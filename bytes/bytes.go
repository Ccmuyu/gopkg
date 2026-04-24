package bytes

import (
	"encoding/hex"
	"encoding/base64"
)

func HexEncode(src []byte) string {
	return hex.EncodeToString(src)
}

func HexDecode(s string) ([]byte, error) {
	return hex.DecodeString(s)
}

func HexEncodeToString(src []byte) string {
	return hex.EncodeToString(src)
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

func HumanReadable(bytes int64) string {
	const unit = 1024
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
	if f < 0 {
		return "-" + floatToStr(-f)
	}
	return formatFloat(f)
}

func formatFloat(f float64) string {
	if f == float64(int64(f)) {
		return intToStr(int64(f))
	}
	var intPart int64 = int64(f)
	var result []byte
	f -= float64(intPart)
	for i := 0; i < 2; i++ {
		f *= 10
		d := byte('0' + int(f)%10)
		f -= float64(int(f))
		result = append(result, d)
	}
	return intToStr(intPart) + "." + string(result)
}