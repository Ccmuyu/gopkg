# logs

轻量级日志库，支持 context 链路追踪、多种日志级别、格式化输出。

## 安装

```go
import "github.com/Ccmuyu/gopkg/logs"
```

## 快速开始

```go
import "context"
"github.com/Ccmuyu/gopkg/logs"

func main() {
    ctx := logs.NewContext(context.Background())
    
    logs.Info(ctx, "server started")
    logs.Warn(ctx, "memory usage high: %d%%", 85)
    logs.Error(ctx, "connection failed: %s", "timeout")
}
```

输出示例：

```
2026-04-24 15:30:00.123 [INFO] [] server started
2026-04-24 15:30:01.456 [WARN] [] memory usage high: 85%
2026-04-24 15:30:02.789 [ERROR] [] connection failed: timeout
```

## Context 链路追踪

使用 `logs.NewContext` 自动生成 trace_id，或 `logs.WithTraceID` 设置指定值：

```go
// 自动生成 trace_id
ctx := logs.NewContext(context.Background())
logs.Info(ctx, "request processed")

// 或指定 trace_id
ctx = logs.WithTraceID(context.Background(), "abc123-456")
logs.Info(ctx, "request processed")

// 输出: 2026-04-24 15:30:00.123 [INFO] [abc123-456] request processed
```

## 自定义 Logger

```go
import (
    "bytes"
    "github.com/Ccmuyu/gopkg/logs"
)

func main() {
    buf := &bytes.Buffer{}
    l := logs.New(
        logs.WithLevel(logs.DEBUG),
        logs.WithOutput(buf),
        logs.WithFormatter(&logs.JSONFormatter{}),
    )
    
    ctx := logs.WithTraceID(context.Background(), "trace-001")
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
l := logs.New(logs.WithFormatter(&logs.JSONFormatter{}))
```

输出：

```json
{"time":"2026-04-24 15:30:00.123","level":"INFO","trace_id":"abc123","msg":"message"}
```

## API 参考

### 全局 Logger

```go
logs.Debug(ctx, format, args...)
logs.Info(ctx, format, args...)
logs.Warn(ctx, format, args...)
logs.Error(ctx, format, args...)
logs.Fatal(ctx, format, args...)
```

### 自定义 Logger

```go
l := logs.New(opts...)
l.Debug(ctx, format, args...)
l.Info(ctx, format, args...)
l.Warn(ctx, format, args...)
l.Error(ctx, format, args...)
l.Fatal(ctx, format, args...)
```

### Options

```go
logs.WithLevel(level)           // 设置日志级别
logs.WithOutput(writer)         // 设置输出目标，默认 os.Stdout
logs.WithFormatter(formatter)   // 设置格式化器
```

### Trace ID

```go
logs.NewTraceID()                      // 生成 128 位随机 trace ID (32 位 hex 字符串)
logs.NewContext(ctx)                    // 生成新 context 并自动注入 trace_id
logs.WithTraceID(ctx, traceID)          // 将指定 trace_id 注入 context
```

## 性能

Benchmark 测试结果（约）：

| 操作 | 耗时 | 内存 |
|------|------|------|
| Info with ctx | ~69μs/op | 247 B/op |