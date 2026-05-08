# bytes

字节操作工具：Hex、Base64 编解码及可读性格式化。

## 安装

```go
import "github.com/Ccmuyu/gopkg/bytes"
```

## 快速开始

```go
bytes.HexEncode([]byte("hello"))                  // 68656c6c6f
bytes.HexDecode("68656c6c6f")                     // []byte("hello")

bytes.Base64Encode([]byte("hello"))                // aGVsbG8=
bytes.Base64Decode("aGVsbG8=")                    // []byte("hello")
bytes.Base64EncodeString("hello")                  // aGVsbG8=
bytes.Base64DecodeString("aGVsbG8=")              // "hello"

bytes.Base64URLEncode([]byte("hello?world="))      // aGVsbG8_d29ybGQ9
bytes.Base64URLEncodeString("hello?world=")        // aGVsbG8_d29ybGQ9

bytes.HumanReadable(1048576)                       // 1 MB
```

## API 参考

```go
func HexEncode(src []byte) string
func HexDecode(s string) ([]byte, error)
func Base64Encode(src []byte) string
func Base64Decode(s string) ([]byte, error)
func Base64EncodeString(s string) string
func Base64DecodeString(s string) (string, error)
func Base64URLEncode(src []byte) string
func Base64URLDecode(s string) ([]byte, error)
func Base64URLEncodeString(s string) string
func Base64URLDecodeString(s string) (string, error)
func HumanReadable(bytes int64) string
```
