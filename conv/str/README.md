# conv/str

字符串与数字互转工具函数。

## 安装

```go
import "github.com/Ccmuyu/gopkg/conv/str"
```

## 快速开始

```go
str.Atoi("123")              // 123
str.Atoi64("123456789")      // 123456789
str.Itoa(123)                // "123"
str.Itoa64(123456789)        // "123456789"
str.ParseFloat("3.14")       // 3.14
str.FormatFloat(3.14)        // "3.14"
str.ParseBool("true")        // true
str.FormatBool(true)         // "true"
```

## API 参考

```go
func Atoi(s string) int
func Atoi64(s string) int64
func Atoi32(s string) int32
func Itoa(i int) string
func Itoa64(i int64) string
func ParseFloat(s string) float64
func FormatFloat(f float64) string
func FormatFloatPrec(f float64, prec int) string
func ParseBool(s string) bool
func FormatBool(b bool) string
```
