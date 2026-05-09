package logs

import (
	"context"
	"crypto/rand"
	"encoding/hex"
)

type contextKey string

const TraceIDKey contextKey = "trace_id"

func NewTraceID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDKey, traceID)
}

func NewContext(ctx context.Context) context.Context {
	return WithTraceID(ctx, NewTraceID())
}
