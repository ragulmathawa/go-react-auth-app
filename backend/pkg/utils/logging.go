package utils

import (
	"log/slog"
	"os"
)

func InitLogging(config AppConfig) {
	var logger *slog.Logger
	if config.Mode == "production" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	} else {
		logger = slog.Default()
	}
	slog.SetDefault(logger)
	slog.SetLogLoggerLevel(getLogLevel(config))
}

func getLogLevel(config AppConfig) slog.Level {
	level := config.LogLevel

	if config.LogLevel == "error" {
		return slog.LevelError
	} else if level == "warn" {
		return slog.LevelWarn
	} else if level == "info" {
		return slog.LevelInfo
	} else if level == "debug" {
		return slog.LevelDebug
	} else {
		return slog.LevelInfo
	}
}
