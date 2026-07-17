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
	mu        sync.Mutex
	level     Level
	output    io.Writer
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
		level:     INFO,
		output:    os.Stdout,
		formatter: &DefaultFormatter{},
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

var defaultLogger atomic.Pointer[Logger]
var once sync.Once

// Default 返回全局默认 Logger。若已通过 SetDefault 设置过则直接返回该实例，
// 否则惰性创建一个默认 Logger。调用顺序无关：先 SetDefault 再 Default 不会被覆盖。
func Default() *Logger {
	if l := defaultLogger.Load(); l != nil {
		return l
	}
	once.Do(func() {
		defaultLogger.CompareAndSwap(nil, New())
	})
	return defaultLogger.Load()
}

// SetDefault 设置全局默认 Logger。可在首次使用包级日志函数（如 logs.Info）之前调用，
// 设置的实例不会被 Default 的惰性初始化覆盖。
func SetDefault(l *Logger) {
	if l == nil {
		return
	}
	defaultLogger.Store(l)
}

func (l *Logger) log(ctx context.Context, level Level, format string, args ...any) {
	if level < l.level {
		return
	}
	msg := fmt.Sprintf(format, args...)
	entry := &Entry{
		Time:  time.Now(),
		Level: level,
		Msg:   msg,
		Ctx:   ctx,
	}
	formatted := l.formatter.Format(entry)
	l.mu.Lock()
	defer l.mu.Unlock()
	_, _ = l.output.Write([]byte(formatted))
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
