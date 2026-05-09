package logs

import (
	"bytes"
	"context"
	"strings"
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func BenchmarkLog(b *testing.B) {
	buf := &bytes.Buffer{}
	l := New(WithOutput(buf), WithLevel(DEBUG))
	ctx := WithTraceID(context.Background(), "test-trace")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Info(ctx, "this is a log message with trace id")
	}
}

func BenchmarkLogWithCtx(b *testing.B) {
	buf := &bytes.Buffer{}
	l := New(WithOutput(buf), WithLevel(DEBUG))
	ctx := WithTraceID(context.Background(), "test-trace")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Info(ctx, "log message: %d", i)
	}
}

func BenchmarkDefaultLogger(b *testing.B) {
	buf := &bytes.Buffer{}
	SetDefault(New(WithOutput(buf), WithLevel(DEBUG)))
	ctx := WithTraceID(context.Background(), "test-trace")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Info(ctx, "log message: %d", i)
	}
}

func TestLogWithTraceID(t *testing.T) {
	buf := &bytes.Buffer{}
	l := New(WithOutput(buf), WithLevel(DEBUG), WithFormatter(&DefaultFormatter{}))

	ctx := WithTraceID(context.Background(), "trace-123")
	l.Info(ctx, "test message")

	result := buf.String()
	AssertTrue(t, strings.Contains(result, "trace-123"))
}

func TestLogWithJSONFormatter(t *testing.T) {
	buf := &bytes.Buffer{}
	l := New(WithOutput(buf), WithLevel(DEBUG), WithFormatter(&JSONFormatter{}))

	ctx := WithTraceID(context.Background(), "trace-456")
	l.Info(ctx, "test message")

	result := buf.String()
	AssertTrue(t, strings.Contains(result, `"trace_id":"trace-456"`))
}