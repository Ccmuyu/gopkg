package log

import (
	"context"
	"testing"
	"time"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestLevel(t *testing.T) {
	AssertEqual(t, DEBUG.String(), "DEBUG")
	AssertEqual(t, INFO.String(), "INFO")
	AssertEqual(t, WARN.String(), "WARN")
	AssertEqual(t, ERROR.String(), "ERROR")
	AssertEqual(t, FATAL.String(), "FATAL")
}

func TestDefaultFormatter(t *testing.T) {
	f := &DefaultFormatter{}
	ctx := context.WithValue(context.Background(), "trace_id", "abc123")
	entry := &Entry{
		Time:  time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
		Level: INFO,
		Msg:   "test message",
		Ctx:   ctx,
	}
	result := f.Format(entry)
	AssertTrue(t, len(result) > 0)
}

func TestJSONFormatter(t *testing.T) {
	f := &JSONFormatter{}
	ctx := context.WithValue(context.Background(), "trace_id", "abc123")
	entry := &Entry{
		Time:  time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
		Level: INFO,
		Msg:   "test message",
		Ctx:   ctx,
	}
	result := f.Format(entry)
	AssertTrue(t, contains(result, "trace_id"))
	AssertTrue(t, contains(result, "abc123"))
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestLoggerLevel(t *testing.T) {
	l := New(WithLevel(INFO))
	l.Debug(nil, "debug")
	l.Info(nil, "info")
}

func TestLoggerOptions(t *testing.T) {
	l := New(
		WithLevel(DEBUG),
		WithFormatter(&JSONFormatter{}),
	)
	AssertTrue(t, l != nil)
}