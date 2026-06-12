# rand

随机工具：随机字符串、整数、字节切片、随机选择、洗牌。

## 安装

```go
import "github.com/Ccmuyu/gopkg/rand"
```

## 快速开始

```go
rand.String(16)                          // "aB3xK9mQ..."
rand.Int(1, 100)                         // 42
rand.Bytes(32)                           // []byte (crypto/rand)
rand.Choice([]string{"a", "b", "c"})     // "a" (随机)
rand.Shuffle([]int{1, 2, 3, 4, 5})       // 原地洗牌
```

## API 参考

```go
func String(n int) string
func Int(min, max int) int
func Bytes(n int) []byte
func Choice[T any](slice []T) T
func Shuffle[T any](slice []T)
```
