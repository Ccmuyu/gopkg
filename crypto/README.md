# crypto

哈希工具函数，基于 Go 标准库实现，无第三方依赖。

## 安装

```go
import "github.com/Ccmuyu/gopkg/crypto"
```

## 快速开始

```go
crypto.MD5("hello")          // 5d41402abc4b2a76b9719d911017c592
crypto.SHA1("hello")         // aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
crypto.SHA256("hello")       // 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
```

## API 参考

```go
func MD5(s string) string
func MD5Bytes(data []byte) string
func SHA1(s string) string
func SHA1Bytes(data []byte) string
func SHA256(s string) string
func SHA256Bytes(data []byte) string
```

## 测试

```bash
go test ./crypto/ -v
```
