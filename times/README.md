# times

时间格式化与解析工具。

## 安装

```go
import "github.com/Ccmuyu/gopkg/times"
```

## 快速开始

```go
t := time.Now()
times.Format(t)                            // "2026-04-24 15:04:05"
times.FormatF(t, "2006/01/02")             // "2026/04/24"
times.ParseUTCTime("2026-04-24T15:04:05Z") // time.Time
```

## API 参考

```go
func Format(t time.Time) string
func FormatF(t time.Time, layout string) string
func ParseUTCTime(t string) time.Time
func ParseUTCTimePtr(t *string) time.Time
```
