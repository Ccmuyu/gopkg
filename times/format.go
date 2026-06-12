package times

import "time"

const (
	layoutDefault = "2006-01-02 15:04:05"
	layoutUTC     = "2006-01-02T15:04:05Z"
	layoutDate    = "2006-01-02"
)

func Format(t time.Time) string {
	return t.Format(layoutDefault)
}

func FormatF(t time.Time, layout string) string {
	return t.Format(layout)
}

func FormatDate(t time.Time) string {
	return t.Format(layoutDate)
}

func ParseUTCTime(t string) time.Time {
	parse, err := time.Parse(layoutUTC, t)
	if err != nil {
		return time.Time{}
	}
	return parse
}

func ParseUTCTimePtr(t *string) time.Time {
	if t == nil {
		return time.Time{}
	}
	return ParseUTCTime(*t)
}

func Unix(sec, nsec int64) time.Time {
	return time.Unix(sec, nsec)
}

func FromUnix(sec int64) time.Time {
	return time.Unix(sec, 0)
}

func BeginOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, 999999999, t.Location())
}

func BeginOfWeek(t time.Time) time.Time {
	weekday := t.Weekday()
	if weekday == time.Sunday {
		weekday = 7
	}
	return BeginOfDay(t.AddDate(0, 0, -int(weekday-time.Monday)))
}

func EndOfWeek(t time.Time) time.Time {
	weekday := t.Weekday()
	if weekday == time.Sunday {
		weekday = 7
	}
	return EndOfDay(t.AddDate(0, 0, 7-int(weekday)))
}

func BeginOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

func EndOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m+1, 0, 23, 59, 59, 999999999, t.Location())
}

func BeginOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

func EndOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 12, 31, 23, 59, 59, 999999999, t.Location())
}

func IsWeekend(t time.Time) bool {
	return t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
}

func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func DaysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func DaysBetween(start, end time.Time) int {
	return int(end.Sub(start).Hours() / 24)
}

func Elapsed(t time.Time) time.Duration {
	return time.Since(t)
}

func Since(t time.Time) time.Duration {
	return time.Since(t)
}

func Until(t time.Time) time.Duration {
	return time.Until(t)
}
