package logs

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
	ctx := WithTraceID(context.Background(), "abc123")
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
	ctx := WithTraceID(context.Background(), "abc123")
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

func TestNewTraceID(t *testing.T) {
	id1 := NewTraceID()
	id2 := NewTraceID()
	AssertTrue(t, len(id1) == 32)
	AssertTrue(t, len(id2) == 32)
	AssertTrue(t, id1 != id2)
}

func TestWithTraceID(t *testing.T) {
	ctx := WithTraceID(context.Background(), "test-trace")
	v := ctx.Value(TraceIDKey)
	AssertEqual(t, v.(string), "test-trace")
}

func TestNewContext(t *testing.T) {
	ctx := NewContext(context.Background())
	v := ctx.Value(TraceIDKey)
	AssertTrue(t, len(v.(string)) == 32)
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