package logs

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

type Logger struct {
	mu      sync.Mutex
	level   Level
	output  io.Writer
	formatter Formatter
}

type Option func(*Logger)

func WithLevel(level Level) Option {
	return func(l *Logger) {
		l.level = level
	}
}

func WithOutput(w io.Writer) Option {
	return func(l *Logger) {
		l.output = w
	}
}

func WithFormatter(f Formatter) Option {
	return func(l *Logger) {
		l.formatter = f
	}
}

func New(opts ...Option) *Logger {
	l := &Logger{
		level:   INFO,
		output:  os.Stdout,
		formatter: &DefaultFormatter{},
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

var defaultLogger atomic.Pointer[Logger]
var once sync.Once

func Default() *Logger {
	once.Do(func() {
		defaultLogger.Store(New())
	})
	return defaultLogger.Load()
}

func SetDefault(l *Logger) {
	defaultLogger.Store(l)
}

func (l *Logger) log(ctx context.Context, level Level, format string, args ...any) {
	if level < l.level {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	msg := fmt.Sprintf(format, args...)
	entry := &Entry{
		Time:  time.Now(),
		Level: level,
		Msg:   msg,
		Ctx:   ctx,
	}
	l.output.Write([]byte(l.formatter.Format(entry)))
}

func (l *Logger) Debug(ctx context.Context, format string, args ...any) {
	l.log(ctx, DEBUG, format, args...)
}

func (l *Logger) Info(ctx context.Context, format string, args ...any) {
	l.log(ctx, INFO, format, args...)
}

func (l *Logger) Warn(ctx context.Context, format string, args ...any) {
	l.log(ctx, WARN, format, args...)
}

func (l *Logger) Error(ctx context.Context, format string, args ...any) {
	l.log(ctx, ERROR, format, args...)
}

func (l *Logger) Fatal(ctx context.Context, format string, args ...any) {
	l.log(ctx, FATAL, format, args...)
	os.Exit(1)
}

func Debug(ctx context.Context, format string, args ...any) {
	Default().Debug(ctx, format, args...)
}

func Info(ctx context.Context, format string, args ...any) {
	Default().Info(ctx, format, args...)
}

func Warn(ctx context.Context, format string, args ...any) {
	Default().Warn(ctx, format, args...)
}

func Error(ctx context.Context, format string, args ...any) {
	Default().Error(ctx, format, args...)
}

func Fatal(ctx context.Context, format string, args ...any) {
	Default().Fatal(ctx, format, args...)
}