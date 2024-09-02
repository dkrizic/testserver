package errorh

import (
	"encoding/json"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func ErrorHandler(next func(w http.ResponseWriter, r *http.Request) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := next(w, r); err != nil {
			message := err.Error()
			ctx := r.Context()
			span := trace.SpanFromContext(ctx)
			span.RecordError(err)
			slog.WarnContext(ctx, "Error handling request", "error", err)
			er := errorResponse{Error: message}
			output, _ := json.Marshal(er)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(output)
		}
	})
}
