package logger

import (
	"log/slog"
	"os"
)

var opts = &slog.HandlerOptions{
	Level: slog.LevelDebug,
}

func SetUpLogger(env string) {

	switch env {
	case "production":
		setUpProductionLogger()
	default:
		setUpDevelopmentLogger()
	}
}

func setUpDevelopmentLogger() {
	logger := slog.New(slog.NewTextHandler(os.Stdin, opts))
	slog.SetDefault(logger)
}

func setUpProductionLogger() {
	opts.AddSource = true
	opts.Level = slog.LevelInfo
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)
}
