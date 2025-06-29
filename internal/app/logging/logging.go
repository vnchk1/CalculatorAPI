package logging

import (
	"github.com/vnchk1/CalculatorAPI/configs"
	"log/slog"
	"os"
)

func HandlerInit() *slog.JSONHandler {
	cfg, err := configs.LoadConfig()
	if err != nil {
		slog.Error("Error loading config", "err", err)
	}
	logLvlStr := cfg.Logger.Level
	logLevel := configs.ConvertLogLevel(logLvlStr)
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	})
	return logHandler
}

func Logger() *slog.Logger {
	logger := slog.New(HandlerInit())
	return logger
}
