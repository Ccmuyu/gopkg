package str

import "strings"

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

func Ellipsis(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
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