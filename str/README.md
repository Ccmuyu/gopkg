# str

字符串操作工具：Trim、Split、Join、Contains、Ellipsis 等。

## 安装

```go
import "github.com/Ccmuyu/gopkg/str"
```

## 快速开始

```go
str.Trim("  hello  ")                      // "hello"
str.Upper("hello")                         // "HELLO"
str.Lower("HELLO")                         // "hello"
str.Split("a,b,c", ",")                    // ["a", "b", "c"]
str.Join([]string{"a", "b", "c"}, ",")     // "a,b,c"
str.Contains("hello world", "world")       // true
str.HasPrefix("hello", "he")               // true
str.HasSuffix("hello", "lo")               // true
str.Replace("hello", "l", "x")             // "hexxo"
str.Ellipsis("hello world", 8)             // "hello..."
str.IsEmpty("")                            // true
str.DefaultIfEmpty("", "default")          // "default"
str.Repeat("ab", 3)                        // "ababab"
str.ContainsAny("hello", "aeiou")          // true
str.Count("hello", "l")                    // 2
```

## API 参考

```go
func Trim(s string) string
func Upper(s string) string
func Lower(s string) string
func Split(s string, sep string) []string
func SplitN(s string, sep string, n int) []string
func Join(arr []string, sep string) string
func Contains(s, substr string) bool
func HasPrefix(s, prefix string) bool
func HasSuffix(s, suffix string) bool
func Replace(s, old, new string) string
func ReplaceN(s, old, new string, n int) string
func Ellipsis(s string, maxLen int) string
func IsEmpty(s string) bool
func DefaultIfEmpty(s string, defaultVal string) string
func Repeat(s string, count int) string
func ContainsAny(s, chars string) bool
func Count(s, sep string) int
```
