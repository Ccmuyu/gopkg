# conv/json

JSON 序列化与反序列化工具函数。

## 安装

```go
import "github.com/Ccmuyu/gopkg/conv/json"
```

## 快速开始

```go
json.Marshal(map[string]int{"a": 1})          // {"a":1}, nil
json.MustMarshal(map[string]int{"a": 1})       // {"a":1}
json.Unmarshal(`{"a":1}`, &v)                 // nil
json.MustUnmarshal(`{"a":1}`, &v)             // 自动忽略 error
json.MarshalIndent(v, "", "  ")               // 带缩进输出
```

## API 参考

```go
func Marshal(v any) ([]byte, error)
func MustMarshal(v any) []byte
func Unmarshal(data string, v any) error
func MustUnmarshal(data string, v any)
func MarshalIndent(v any, prefix, indent string) ([]byte, error)
```
