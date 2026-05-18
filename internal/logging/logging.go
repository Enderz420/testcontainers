package logging

import (
	"context"
	"log/slog"
)


type ContextKey string

const LoggerKey ContextKey = "logger"

func LoggerFromContext(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(LoggerKey).(*slog.Logger)
	if !ok {
		return slog.Default()
	}
	return logger
}