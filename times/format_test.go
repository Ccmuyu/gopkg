package times

import (
	"testing"
	"time"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestFormat(t *testing.T) {
	unix := time.Unix(1726675518, 0)
	AssertEqual(t, Format(unix), "2024-09-19 00:05:18")
	AssertNotEqual(t, Format(unix), "=2024-09-19 00:05:18")
}

func TestFormatF(t *testing.T) {
	unix := time.Unix(1726675518, 0)
	AssertEqual(t, FormatF(unix, layoutUTC), "2024-09-19T00:05:18Z")
	AssertNotEqual(t, FormatF(unix, layoutUTC), "~2024-09-19T00:05:18Z")
}

func TestFormatDate(t *testing.T) {
	unix := time.Unix(1726675518, 0)
	AssertEqual(t, FormatDate(unix), "2024-09-19")
}

func TestParseUTCTime(t *testing.T) {
	tm := ParseUTCTime("2024-09-19T00:05:18Z")
	AssertEqual(t, tm.Year(), 2024)
	AssertEqual(t, tm.Month(), time.September)
	AssertEqual(t, tm.Day(), 19)
	AssertEqual(t, tm.Hour(), 0)
	AssertEqual(t, tm.Minute(), 5)
	AssertEqual(t, tm.Second(), 18)
	AssertTrue(t, ParseUTCTime("invalid").IsZero())
}

func TestParseUTCTimePtr(t *testing.T) {
	s := "2024-09-19T00:05:18Z"
	tm := ParseUTCTimePtr(&s)
	AssertEqual(t, tm.Year(), 2024)
	AssertEqual(t, tm.Month(), time.September)
	AssertEqual(t, tm.Day(), 19)
	AssertTrue(t, ParseUTCTimePtr(nil).IsZero())
}

func TestUnix(t *testing.T) {
	AssertEqual(t, Unix(1726675518, 0).Unix(), int64(1726675518))
}

func TestFromUnix(t *testing.T) {
	AssertEqual(t, FromUnix(1726675518).Unix(), int64(1726675518))
}

func TestBeginOfDay(t *testing.T) {
	now := time.Date(2024, 9, 19, 15, 30, 45, 123, time.UTC)
	start := BeginOfDay(now)
	AssertEqual(t, start.Year(), 2024)
	AssertEqual(t, start.Month(), time.September)
	AssertEqual(t, start.Day(), 19)
	AssertEqual(t, start.Hour(), 0)
	AssertEqual(t, start.Minute(), 0)
	AssertEqual(t, start.Second(), 0)
	AssertEqual(t, start.Nanosecond(), 0)
}

func TestEndOfDay(t *testing.T) {
	now := time.Date(2024, 9, 19, 15, 30, 45, 123, time.UTC)
	end := EndOfDay(now)
	AssertEqual(t, end.Year(), 2024)
	AssertEqual(t, end.Month(), time.September)
	AssertEqual(t, end.Day(), 19)
	AssertEqual(t, end.Hour(), 23)
	AssertEqual(t, end.Minute(), 59)
	AssertEqual(t, end.Second(), 59)
}

func TestBeginOfWeek(t *testing.T) {
	wednesday := time.Date(2024, 9, 18, 15, 0, 0, 0, time.UTC)
	monday := BeginOfWeek(wednesday)
	AssertEqual(t, monday.Weekday(), time.Monday)
	AssertEqual(t, monday.Day(), 16)
	AssertEqual(t, monday.Hour(), 0)
}

func TestEndOfWeek(t *testing.T) {
	wednesday := time.Date(2024, 9, 18, 15, 0, 0, 0, time.UTC)
	sunday := EndOfWeek(wednesday)
	AssertEqual(t, sunday.Weekday(), time.Sunday)
	AssertEqual(t, sunday.Day(), 22)
	AssertEqual(t, sunday.Hour(), 23)
	AssertEqual(t, sunday.Minute(), 59)
	AssertEqual(t, sunday.Second(), 59)
}

func TestEndOfWeekSunday(t *testing.T) {
	sun := time.Date(2024, 9, 22, 12, 0, 0, 0, time.UTC)
	end := EndOfWeek(sun)
	AssertEqual(t, end.Day(), 22)
	AssertEqual(t, end.Hour(), 23)
}

func TestBeginOfWeekSunday(t *testing.T) {
	sun := time.Date(2024, 9, 22, 12, 0, 0, 0, time.UTC)
	mon := BeginOfWeek(sun)
	AssertEqual(t, mon.Weekday(), time.Monday)
	AssertEqual(t, mon.Day(), 16)
}

func TestBeginOfMonth(t *testing.T) {
	now := time.Date(2024, 9, 19, 15, 30, 0, 0, time.UTC)
	start := BeginOfMonth(now)
	AssertEqual(t, start.Day(), 1)
	AssertEqual(t, start.Month(), time.September)
	AssertEqual(t, start.Hour(), 0)
}

func TestEndOfMonth(t *testing.T) {
	now := time.Date(2024, 9, 19, 15, 30, 0, 0, time.UTC)
	end := EndOfMonth(now)
	AssertEqual(t, end.Day(), 30)
	AssertEqual(t, end.Month(), time.September)
	AssertEqual(t, end.Hour(), 23)
}

func TestEndOfMonthFebLeap(t *testing.T) {
	now := time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC)
	end := EndOfMonth(now)
	AssertEqual(t, end.Day(), 29)
}

func TestBeginOfYear(t *testing.T) {
	now := time.Date(2024, 9, 19, 15, 30, 0, 0, time.UTC)
	start := BeginOfYear(now)
	AssertEqual(t, start.Month(), time.January)
	AssertEqual(t, start.Day(), 1)
	AssertEqual(t, start.Hour(), 0)
}

func TestEndOfYear(t *testing.T) {
	now := time.Date(2024, 9, 19, 15, 30, 0, 0, time.UTC)
	end := EndOfYear(now)
	AssertEqual(t, end.Month(), time.December)
	AssertEqual(t, end.Day(), 31)
	AssertEqual(t, end.Hour(), 23)
}

func TestIsWeekend(t *testing.T) {
	saturday := time.Date(2024, 9, 21, 0, 0, 0, 0, time.UTC)
	sunday := time.Date(2024, 9, 22, 0, 0, 0, 0, time.UTC)
	monday := time.Date(2024, 9, 23, 0, 0, 0, 0, time.UTC)
	AssertTrue(t, IsWeekend(saturday))
	AssertTrue(t, IsWeekend(sunday))
	AssertTrue(t, !IsWeekend(monday))
}

func TestIsLeapYear(t *testing.T) {
	AssertTrue(t, IsLeapYear(2024))
	AssertTrue(t, IsLeapYear(2000))
	AssertTrue(t, !IsLeapYear(1900))
	AssertTrue(t, !IsLeapYear(2023))
}

func TestDaysInMonth(t *testing.T) {
	AssertEqual(t, DaysInMonth(2024, time.January), 31)
	AssertEqual(t, DaysInMonth(2024, time.February), 29)
	AssertEqual(t, DaysInMonth(2023, time.February), 28)
	AssertEqual(t, DaysInMonth(2024, time.April), 30)
}

func TestDaysBetween(t *testing.T) {
	start := time.Date(2024, 9, 19, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 9, 25, 0, 0, 0, 0, time.UTC)
	AssertEqual(t, DaysBetween(start, end), 6)
	AssertEqual(t, DaysBetween(end, start), -6)
	AssertEqual(t, DaysBetween(start, start), 0)
}
