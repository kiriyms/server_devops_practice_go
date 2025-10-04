package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

func Logger(next http.Handler) http.Handler {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Info("request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.String("remote", r.RemoteAddr),
			slog.Duration("duration", time.Since(start)),
		)
	})
}
