package logs

import (
	"context"
	"io"
	"sync"
	"sync/atomic"
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

// TestSetDefaultBeforeDefault 保证在首次 Default() 之前调用 SetDefault
// 不会被惰性初始化覆盖。
func TestSetDefaultBeforeDefault(t *testing.T) {
	resetDefaultState()

	custom := New(WithLevel(DEBUG))
	SetDefault(custom)
	AssertTrue(t, Default() == custom)
}

// TestDefaultLazyInit 在未设置时 Default() 惰性创建非 nil logger。
func TestDefaultLazyInit(t *testing.T) {
	resetDefaultState()

	l := Default()
	AssertNotNil(t, l)
	// 再次调用返回同一实例
	AssertTrue(t, Default() == l)
}

func TestSetDefaultIgnoresNil(t *testing.T) {
	resetDefaultState()
	original := Default()
	SetDefault(nil)
	AssertTrue(t, Default() == original)
}

type reentrantFormatter struct {
	logger  *Logger
	entered atomic.Bool
}

func (f *reentrantFormatter) Format(*Entry) string {
	if f.entered.CompareAndSwap(false, true) {
		f.logger.Info(context.Background(), "nested")
	}
	return "formatted\n"
}

func TestFormatterCanLogWithoutDeadlock(t *testing.T) {
	formatter := &reentrantFormatter{}
	logger := New(WithOutput(io.Discard), WithFormatter(formatter))
	formatter.logger = logger

	done := make(chan struct{})
	go func() {
		logger.Info(context.Background(), "outer")
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("logging deadlocked in formatter")
	}
}

func resetDefaultState() {
	once = sync.Once{}
	defaultLogger.Store(nil)
}
