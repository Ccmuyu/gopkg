# times

时间格式化、解析与计算工具。

## 安装

```go
import "github.com/Ccmuyu/gopkg/times"
```

## 快速开始

```go
t := time.Now()
times.Format(t)                              // "2026-04-24 15:04:05"
times.FormatDate(t)                          // "2026-04-24"
times.BeginOfDay(t)                          // 当天 00:00:00
times.EndOfDay(t)                            // 当天 23:59:59
times.BeginOfWeek(t)                         // 当周周一
times.BeginOfMonth(t)                        // 当月1号
times.EndOfMonth(t)                          // 当月最后一天
times.IsWeekend(t)                           // 是否周末
times.DaysBetween(t, future)                 // 相差天数
times.FromUnix(1712345678)                   // 时间戳转时间
```

## API 参考

```go
func Format(t time.Time) string
func FormatF(t time.Time, layout string) string
func FormatDate(t time.Time) string
func ParseUTCTime(t string) time.Time
func ParseUTCTimePtr(t *string) time.Time
func Unix(sec, nsec int64) time.Time
func FromUnix(sec int64) time.Time
func BeginOfDay(t time.Time) time.Time
func EndOfDay(t time.Time) time.Time
func BeginOfWeek(t time.Time) time.Time
func EndOfWeek(t time.Time) time.Time
func BeginOfMonth(t time.Time) time.Time
func EndOfMonth(t time.Time) time.Time
func BeginOfYear(t time.Time) time.Time
func EndOfYear(t time.Time) time.Time
func IsWeekend(t time.Time) bool
func IsLeapYear(year int) bool
func DaysInMonth(year int, month time.Month) int
func DaysBetween(start, end time.Time) int
func Elapsed(t time.Time) time.Duration
func Since(t time.Time) time.Duration
func Until(t time.Time) time.Duration
```
