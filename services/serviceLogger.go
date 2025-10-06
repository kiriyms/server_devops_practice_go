package services

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) Greet(ctx context.Context) (msg string, err error) {
	defer func(start time.Time) {
		slog.Info(
			"Greet called",
			slog.String("res", msg),
			slog.String("err", fmt.Sprintf("%v", err)),
			slog.Duration("took", time.Since(start)),
		)
	}(time.Now())

	return s.next.Greet(ctx)
}
