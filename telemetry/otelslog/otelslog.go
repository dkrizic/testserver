// Package otelslog provides an [slog.Handler] that attaches OpenTelemetry trace details to logs.
package otelslog

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel/codes"
	"log/slog"

	"go.opentelemetry.io/otel/trace"
)

const (
	mdcKey     = "mdc"
	traceIDKey = "trace_id"
	spanIDKey  = "span_id"
)

// NewHandler returns a new [Handler].
func NewHandler(handler slog.Handler) slog.Handler {
	return Handler{
		Handler: handler,
	}
}

// Middleware returns a [Middleware] for an [slogmulti.Pipe] handler.
//
// [Middleware]: https://pkg.go.dev/github.com/samber/slog-multi#Middleware
// [slogmulti.Pipe]: https://pkg.go.dev/github.com/samber/slog-multi#Pipe
func Middleware() func(slog.Handler) slog.Handler {
	return func(handler slog.Handler) slog.Handler {
		return NewHandler(handler)
	}
}

// Handler attaches details from an OpenTelemetry trace to each log record.
type Handler struct {
	slog.Handler
}

// Handle implements [slog.Handler].
func (h Handler) Handle(ctx context.Context, record slog.Record) error {
	if h.Handler == nil {
		return errors.New("otelslog: handler is missing")
	}

	spanContext := trace.SpanContextFromContext(ctx)
	span := trace.SpanFromContext(ctx)

	if spanContext.HasTraceID() && spanContext.HasSpanID() {
		traceId := spanContext.TraceID().String()
		spanId := spanContext.SpanID().String()
		// add a mdc field and add trace and span id to it
		record.Add(mdcKey, map[string]string{
			traceIDKey: traceId,
			spanIDKey:  spanId,
		})
	}

	if record.Level == slog.LevelError {
		record.Attrs(func(a slog.Attr) bool {
			if a.Key == "error" {
				span.RecordError(errors.New(a.Value.String()))
				span.SetStatus(codes.Error, a.Value.String())
				return false
			}
			return true
		})
	}

	return h.Handler.Handle(ctx, record)
}

// WithAttrs implements [slog.Handler].
func (h Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if h.Handler == nil {
		return h
	}

	return Handler{h.Handler.WithAttrs(attrs)}
}

// WithGroup implements [slog.Handler].
func (h Handler) WithGroup(name string) slog.Handler {
	if h.Handler == nil {
		return h
	}

	return Handler{h.Handler.WithGroup(name)}
}
