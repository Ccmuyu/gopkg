# str

字符串操作工具：Trim、Split、Join、Contains、Ellipsis、大小写转换、脱敏等。

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
str.Ellipsis("hello world", 8)             // "hello wo..."

str.After("user@example.com", "@")         // "example.com"
str.Before("user@example.com", "@")        // "user"
str.Between("{{hello}}", "{{", "}}")       // "hello"
str.Reverse("hello")                       // "olleh"
str.PadLeft("42", 5, "0")                  // "00042"
str.PadRight("42", 5, "0")                 // "42000"
str.Capitalize("hello")                    // "Hello"
str.ToCamel("hello_world")                 // "helloWorld"
str.ToSnake("helloWorld")                  // "hello_world"
str.ToKebab("helloWorld")                  // "hello-world"
str.Mask("13812345678", 3, '*')            // "138********"
str.Truncate("hello world", 5)             // "hello"
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
func After(s, sep string) string
func Before(s, sep string) string
func Between(s, open, close string) string
func Reverse(s string) string
func PadLeft(s string, length int, pad string) string
func PadRight(s string, length int, pad string) string
func Truncate(s string, maxLen int) string
func Capitalize(s string) string
func ToCamel(s string) string
func ToSnake(s string) string
func ToKebab(s string) string
func Mask(s string, unmaskLen int, mask rune) string
```

`Ellipsis` 的 `maxLen` 表示保留的字符数，不包含追加的 `...`；负数按 0 处理。
`Mask` 的负数 `unmaskLen` 同样按 0 处理。
