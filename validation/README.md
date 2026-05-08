# validation

校验工具：邮箱、手机号、URL、IP、字符串格式等验证。

## 安装

```go
import "github.com/Ccmuyu/gopkg/validation"
```

## 快速开始

```go
validation.IsEmail("test@example.com")      // true
validation.IsPhone("13812345678")           // true
validation.IsURL("https://example.com")     // true
validation.IsIP("192.168.1.1")              // true
validation.IsIPv4("192.168.1.1")            // true
validation.IsAlpha("hello")                 // true
validation.IsNumeric("12345")               // true
validation.IsAlphanumeric("abc123")         // true
validation.IsMatch("hello123", `\d+`)       // true
validation.IsEmailAddr("user@example.com")  // true
validation.IsPrivateIP("10.0.0.1")          // true
```

## API 参考

```go
func IsEmail(s string) bool
func IsPhone(s string) bool
func IsURL(s string) bool
func IsIP(s string) bool
func IsIPv4(s string) bool
func IsAlpha(s string) bool
func IsNumeric(s string) bool
func IsAlphanumeric(s string) bool
func IsMatch(s string, pattern string) bool
func IsEmailAddr(s string) bool
func IsPrivateIP(s string) bool
```
