# gopkg

Go 常用工具集，为 Go 开发提供便捷的工具函数。

## 安装

```bash
go get github.com/Ccmuyu/gopkg
```

## 工具包概览

| 包 | 说明 |
|---|---|
| `conv` | 类型转换：指针、字符串、JSON |
| `slices` | 切片操作：过滤、映射、去重、分组等 |
| `maps` | Map 操作：Keys、Values、Merge、MergeWith、Filter、Map、MapKey、Pick、Omit、Invert、ForEach、IsEqual 等 |
| `str` | 字符串工具：Trim、Split、Contains、Ellipsis 等 |
| `errors` | 错误处理：Wrap、Is、As、Join |
| `validation` | 校验工具：IsEmail、IsPhone、IsURL、IsIP 等 |
| `net/ip` | IP 处理：IsValid、IsPrivate、ToInt |
| `net/url` | URL 处理：Parse、ParseQuery、Build |
| `bytes` | 字节操作：Hex、Base64、HumanReadable，含 string 便捷变体 |
| `crypto` | 加密哈希：MD5、SHA1、SHA256 |
| `log` | 日志库：支持 context、多种级别、Formatter |
| `sorts` | 排序：Sort、SortFunc、SortFuncStable 支持自定义比较器 |
| `times` | 时间工具：Format、FormatF、ParseUTCTime |
| `test` | 测试工具：AssertEqual、AssertSliceEqual |

## 快速开始

### 类型转换

```go
import (
    "github.com/Ccmuyu/gopkg/conv/str"
    "github.com/Ccmuyu/gopkg/conv/json"
)

str.Atoi("123")           // 123
str.Itoa(123)             // "123"
json.MustMarshal(map[string]int{"a": 1})  // {"a":1}
```

### 切片操作

```go
import "github.com/Ccmuyu/gopkg/slices"

nums := []int{1, 2, 3, 4, 5}
evens := slices.Filter(nums, func(v int) bool { return v%2 == 0 })
// [2, 4]

doubled := slices.Map(nums, func(v int) int { return v * 2 })
// [2, 4, 6, 8, 10]

slices.Dedup([]int{1, 1, 2, 2, 3})
// [1, 2, 3]
```

### 日志

```go
import "github.com/Ccmuyu/gopkg/log"

ctx := context.WithValue(context.Background(), "trace_id", "abc123")
log.Info(ctx, "request processed")
// 2026-04-24 16:30:00.123 [INFO] [abc123] request processed
```

### 校验

```go
import "github.com/Ccmuyu/gopkg/validation"

validation.IsEmail("test@example.com")  // true
validation.IsPhone("13812345678")       // true
validation.IsIP("192.168.1.1")          // true
```

## 测试

```bash
go test ./... -v
```

## 性能测试

```bash
go test ./... -bench=. -benchmem
```

详细性能报告见 [docs/benchmark-report.md](docs/benchmark-report.md)。

## 规范

1. 使用 Go 1.18+ 泛型
2. 函数简短、直接
3. 每个功能配有测试
4. 测试辅助函数统一在 `test` 包