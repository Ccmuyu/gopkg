package log

import (
	"context"
	"time"
)

type Entry struct {
	Time  time.Time
	Level Level
	Msg   string
	Ctx   context.Context
}

type Formatter interface {
	Format(*Entry) string
}

type DefaultFormatter struct{}

func (f *DefaultFormatter) Format(entry *Entry) string {
	return formatEntry(entry)
}

type JSONFormatter struct{}

func (f *JSONFormatter) Format(entry *Entry) string {
	return formatJSONEntry(entry)
}

func formatEntry(entry *Entry) string {
	traceID := ""
	if entry.Ctx != nil {
		if v := entry.Ctx.Value("trace_id"); v != nil {
			traceID = v.(string)
		}
	}
	return formatTime(entry.Time) + " [" + entry.Level.String() + "] " +
		"[" + traceID + "] " + entry.Msg + "\n"
}

func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05.000")
}

func formatJSONEntry(entry *Entry) string {
	traceID := ""
	if entry.Ctx != nil {
		if v := entry.Ctx.Value("trace_id"); v != nil {
			traceID = v.(string)
		}
	}
	return `{"time":"` + formatTime(entry.Time) + `","level":"` +
		entry.Level.String() + `","trace_id":"` + traceID +
		`","msg":"` + entry.Msg + `"` + "}\n"
}