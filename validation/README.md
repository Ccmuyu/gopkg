# validation

校验工具：邮箱、手机号、URL、IP、UUID、信用卡号、JSON 等验证。

## 安装

```go
import "github.com/Ccmuyu/gopkg/validation"
```

## 快速开始

```go
validation.IsEmail("test@example.com")           // true
validation.IsPhone("13812345678")                // true
validation.IsURL("https://example.com")          // true
validation.IsIP("192.168.1.1")                   // true
validation.IsUUID("550e8400-e29b-41d4-a716-446655440000") // true
validation.IsCreditCard("4111111111111111")       // true
validation.IsJSON(`{"key":"value"}`)             // true
validation.IsMAC("00:1A:2B:3C:4D:5E")            // true
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
func IsUUID(s string) bool
func IsMAC(s string) bool
func IsJSON(s string) bool
func IsHexColor(s string) bool
func IsPostalCode(s string) bool
func IsPort(port int) bool
func IsLatitude(lat float64) bool
func IsLongitude(lng float64) bool
func IsCreditCard(s string) bool
```
