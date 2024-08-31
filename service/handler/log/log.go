package log

import (
	"log/slog"
	"net/http"
)

func LogHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/health" {
			slog.DebugContext(r.Context(), "Request", "method", r.Method, "url", r.URL.String())
		}
		handler.ServeHTTP(w, r)
	})
}
