package config

import (
	"log/slog"
	"os"
	"strings"
)

// figuring out logger level based on env variable LOG_LEVEL otherwise default to info
func getLogLevel() slog.Level {
	logLevel := slog.LevelInfo
	if level := os.Getenv("LOG_LEVEL"); len(level) != 0 {
		switch strings.ToLower(level) {
		case "info":
			logLevel = slog.LevelInfo
			break
		case "debug":
			logLevel = slog.LevelDebug
			break
		case "warn":
			logLevel = slog.LevelWarn
			break
		case "error":
			logLevel = slog.LevelError
			break
		}
	}
	return logLevel
}
func NewLogger() *slog.Logger {
	level := getLogLevel()

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})
	return slog.New(logHandler)
}
