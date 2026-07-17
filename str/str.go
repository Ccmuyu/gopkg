package str

import (
	"strings"
	"unicode"
)

func Trim(s string) string {
	return strings.TrimSpace(s)
}

func Upper(s string) string {
	return strings.ToUpper(s)
}

func Lower(s string) string {
	return strings.ToLower(s)
}

func Split(s string, sep string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, sep)
}

func SplitN(s string, sep string, n int) []string {
	if s == "" {
		return []string{}
	}
	return strings.SplitN(s, sep, n)
}

func Join(arr []string, sep string) string {
	return strings.Join(arr, sep)
}

func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

func Replace(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

func ReplaceN(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// Ellipsis 最多保留 maxLen 个字符，并在发生截断时追加 "...";
// maxLen 不包含省略号长度，负数按 0 处理。
func Ellipsis(s string, maxLen int) string {
	if maxLen < 0 {
		maxLen = 0
	}
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	return string(runes[:maxLen]) + "..."
}

func IsEmpty(s string) bool {
	return s == ""
}

func DefaultIfEmpty(s string, defaultVal string) string {
	if s == "" {
		return defaultVal
	}
	return s
}

func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

func ContainsAny(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}

func Count(s, sep string) int {
	return strings.Count(s, sep)
}

func After(s, sep string) string {
	idx := strings.Index(s, sep)
	if idx < 0 {
		return ""
	}
	return s[idx+len(sep):]
}

func Before(s, sep string) string {
	idx := strings.Index(s, sep)
	if idx < 0 {
		return ""
	}
	return s[:idx]
}

func Between(s, open, close string) string {
	start := strings.Index(s, open)
	if start < 0 {
		return ""
	}
	start += len(open)
	end := strings.Index(s[start:], close)
	if end < 0 {
		return ""
	}
	return s[start : start+end]
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func PadLeft(s string, length int, pad string) string {
	runes := []rune(s)
	if len(runes) >= length {
		return s
	}
	return strings.Repeat(pad, length-len(runes)) + s
}

func PadRight(s string, length int, pad string) string {
	runes := []rune(s)
	if len(runes) >= length {
		return s
	}
	return s + strings.Repeat(pad, length-len(runes))
}

func Truncate(s string, maxLen int) string {
	if maxLen < 0 {
		return s
	}
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	return string(runes[:maxLen])
}

func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

func ToCamel(s string) string {
	words := splitWords(s)
	for i, w := range words {
		if i == 0 {
			words[i] = strings.ToLower(w)
		} else {
			words[i] = capitalizeWord(w)
		}
	}
	return strings.Join(words, "")
}

func ToSnake(s string) string {
	words := splitWords(s)
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return strings.Join(words, "_")
}

func ToKebab(s string) string {
	words := splitWords(s)
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return strings.Join(words, "-")
}

func Mask(s string, unmaskLen int, mask rune) string {
	if unmaskLen < 0 {
		unmaskLen = 0
	}
	masked := []rune(s)
	if len(masked) <= unmaskLen {
		return s
	}
	for i := unmaskLen; i < len(masked); i++ {
		masked[i] = mask
	}
	return string(masked)
}

func splitWords(s string) []string {
	runes := []rune(s)
	var words []string
	var cur []rune
	for i, r := range runes {
		if unicode.IsUpper(r) && i > 0 && (unicode.IsLower(runes[i-1]) || (i+1 < len(runes) && unicode.IsLower(runes[i+1]))) {
			if len(cur) > 0 {
				words = append(words, string(cur))
				cur = nil
			}
		}
		if r == '_' || r == '-' || r == ' ' {
			if len(cur) > 0 {
				words = append(words, string(cur))
				cur = nil
			}
			continue
		}
		cur = append(cur, r)
	}
	if len(cur) > 0 {
		words = append(words, string(cur))
	}
	return words
}

func capitalizeWord(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}
