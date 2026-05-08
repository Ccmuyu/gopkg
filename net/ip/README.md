# net/ip

IP 地址验证与转换工具。

## 安装

```go
import "github.com/Ccmuyu/gopkg/net/ip"
```

## 快速开始

```go
ip.IsValid("192.168.1.1")        // true
ip.IsPrivate("10.0.0.1")         // true
ip.ToInt("192.168.1.1")          // 3232235777, nil
ip.IntToIP(3232235777)           // "192.168.1.1"
ip.ToJSON("192.168.1.1")         // {"ip":"192.168.1.1"}, nil
```

## API 参考

```go
func IsValid(s string) bool
func IsPrivate(s string) bool
func ToInt(s string) (int64, error)
func IntToIP(n int64) string
func ToJSON(s string) ([]byte, error)
```
