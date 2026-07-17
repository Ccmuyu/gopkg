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
crypto.RandomToken(32)       // URL-safe Base64 随机令牌
```

MD5、SHA-1 和 CRC32 仅适用于兼容、校验等非安全场景，不能用于密码存储、签名或防篡改。消息认证应使用 HMAC；密码存储应使用专用 KDF。

## API 参考

```go
func MD5(s string) string
func MD5Bytes(data []byte) string
func SHA1(s string) string
func SHA1Bytes(data []byte) string
func SHA256(s string) string
func SHA256Bytes(data []byte) string
func SHA512(s string) string
func SHA512Bytes(data []byte) string
func HMAC(algorithm func() hash.Hash, key, data []byte) string
func CRC32(data []byte) uint32
func CRC32String(s string) uint32
func RandomToken(n int) string
```

`RandomToken` 在 `n <= 0` 或系统随机源失败时返回空字符串。

## 测试

```bash
go test ./crypto/ -v
```
