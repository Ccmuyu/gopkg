# coalesce

从一组值中返回第一个非零值，常用于设置默认值。

## 安装

```go
import "github.com/Ccmuyu/gopkg/coalesce"
```

## 快速开始

```go
coalesce.Coalesce("", "", "hello", "world")  // "hello"
coalesce.Coalesce(0, 0, 42, 100)             // 42
coalesce.Coalesce(0, 0, 0)                   // 0
```

## API 参考

```go
func Coalesce[T comparable](vals ...T) T
func CoalesceSlice[T comparable](vals []T) T
```
