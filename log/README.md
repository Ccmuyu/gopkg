# log

轻量级日志库，支持 context 链路追踪、多种日志级别、格式化输出。

## 安装

```go
import "github.com/Ccmuyu/gopkg/log"
```

## 快速开始

```go
import "context"
"github.com/Ccmuyu/gopkg/log"

func main() {
    ctx := context.Background()
    
    log.Info(ctx, "server started")
    log.Warn(ctx, "memory usage high: %d%%", 85)
    log.Error(ctx, "connection failed: %s", "timeout")
}
```

输出示例：

```
2026-04-24 15:30:00.123 [INFO] [] server started
2026-04-24 15:30:01.456 [WARN] [] memory usage high: 85%
2026-04-24 15:30:02.789 [ERROR] [] connection failed: timeout
```

## Context 链路追踪

在 context 中存入 `trace_id`，日志会自动提取并输出：

```go
ctx := context.WithValue(context.Background(), "trace_id", "abc123-456")
log.Info(ctx, "request processed")

// 输出: 2026-04-24 15:30:00.123 [INFO] [abc123-456] request processed
```

## 自定义 Logger

```go
import (
    "bytes"
    "github.com/Ccmuyu/gopkg/log"
)

func main() {
    buf := &bytes.Buffer{}
    l := log.New(
        log.WithLevel(log.DEBUG),
        log.WithOutput(buf),
        log.WithFormatter(&log.JSONFormatter{}),
    )
    
    ctx := context.WithValue(context.Background(), "trace_id", "trace-001")
    l.Info(ctx, "custom logger test")
}
```

## 日志级别

| 级别 | 说明 |
|------|------|
| DEBUG | 调试信息 |
| INFO  | 一般信息 |
| WARN  | 警告信息 |
| ERROR | 错误信息 |
| FATAL | 致命错误，输出后程序退出 |

## Formatter

### DefaultFormatter（默认）

文本格式：

```
2026-04-24 15:30:00.123 [INFO] [trace_id] message
```

### JSONFormatter

JSON 格式，便于日志收集系统处理：

```go
l := log.New(log.WithFormatter(&log.JSONFormatter{}))
```

输出：

```json
{"time":"2026-04-24 15:30:00.123","level":"INFO","trace_id":"abc123","msg":"message"}
```

## API 参考

### 全局 Logger

```go
log.Debug(ctx, format, args...)
log.Info(ctx, format, args...)
log.Warn(ctx, format, args...)
log.Error(ctx, format, args...)
log.Fatal(ctx, format, args...)
```

### 自定义 Logger

```go
l := log.New(opts...)
l.Debug(ctx, format, args...)
l.Info(ctx, format, args...)
l.Warn(ctx, format, args...)
l.Error(ctx, format, args...)
l.Fatal(ctx, format, args...)
```

### Options

```go
log.WithLevel(level)           // 设置日志级别
log.WithOutput(writer)         // 设置输出目标，默认 os.Stdout
log.WithFormatter(formatter)   // 设置格式化器
```

## 性能

Benchmark 测试结果（约）：

| 操作 | 耗时 | 内存 |
|------|------|------|
| Info with ctx | ~69μs/op | 247 B/op |