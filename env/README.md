# env

环境变量工具：读取、设置、查询环境变量。

## 安装

```go
import "github.com/Ccmuyu/gopkg/env"
```

## 快速开始

```go
env.Get("HOME", "/tmp")           // 读取环境变量，支持默认值
env.MustGet("PATH")               // 读取环境变量，不存在则 panic
env.Set("MY_KEY", "my_value")     // 设置环境变量
env.Has("MY_KEY")                 // true
env.EnvironMap()                  // 所有环境变量 map
```

## API 参考

```go
func Get(key string, defaultVal string) string
func MustGet(key string) string
func Set(key, value string) error
func Unset(key string) error
func Has(key string) bool
func EnvironMap() map[string]string
```
