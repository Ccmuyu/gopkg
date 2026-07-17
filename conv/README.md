# conv

类型转换工具：指针安全取值、字符串与数字互转、JSON 序列化。

## 子包

| 子包 | 说明 |
|------|------|
| `conv` | 指针安全转换（String、Bool、Int、Int64、Int32 等） |
| `conv/str` | 字符串与数字互转（Atoi、Itoa、ParseFloat、FormatFloat 等） |
| `conv/json` | JSON 序列化与反序列化（Marshal、Unmarshal、MustMarshal 等） |

## 安装

```go
import "github.com/Ccmuyu/gopkg/conv"
import "github.com/Ccmuyu/gopkg/conv/str"
import "github.com/Ccmuyu/gopkg/conv/json"
```

## 快速开始

```go
// 指针安全取值（nil 返回零值）
conv.String(ptrStr)          // *string -> string
conv.Int(ptrInt)             // *int -> int
conv.Int64Ptr(42)            // int -> *int64

// 字符串与数字互转
str.Atoi("123")              // 123
str.Itoa(123)                // "123"
str.ParseFloat("3.14")       // 3.14

// JSON 序列化
json.MustMarshal(map[string]int{"a": 1})  // {"a":1}
```

## API 参考

### conv（指针转换）

```go
func Val[T any](p *T) T
func Ptr[T any](v T) *T
func String(v *string) string
func StringPtr(v string) *string
func Bool(v *bool) bool
func BoolPtr(v bool) *bool
func Int(v *int) int
func IntPtr(v int) *int
func Int16(v *int16) int16
func Int16Ptr(v int16) *int16
func Int32(v *int32) int32
func Int32Ptr(v int32) *int32
func Int64(v *int64) int64
func Int64Ptr(v int64) *int64
func IntPtrToInt64Ptr(v *int) *int64
func Int64PtrToIntPtr(v *int64) *int
func Int64PtrToInt(v *int64) int
```
