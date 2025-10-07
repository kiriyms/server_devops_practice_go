package handlers

import (
	"log/slog"
	"net/http"
	"time"
)

type LoggingHandler struct {
	next http.Handler
}

func NewLoggingHandler(next http.Handler) http.Handler {
	return &LoggingHandler{
		next: next,
	}
}

func (s *LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func(start time.Time) {
		slog.Info(
			"HTTP request",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()),
			slog.Duration("took", time.Since(start)),
		)
	}(time.Now())

	s.next.ServeHTTP(w, r)
}
