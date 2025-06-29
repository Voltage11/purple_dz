package logger

import (
	"log/slog"
	"os"
	"purple1/config"
)

func NewLogger(config *config.LogConfig) *slog.Logger {
	slog.SetLogLoggerLevel(slog.Level(config.Level))
	
	var logger slog.Logger
	if config.Format == "json" {
		logger = *slog.New(slog.NewJSONHandler(os.Stdout, nil))
	} else {
		logger = *slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	return &logger
}