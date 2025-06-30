package logging

import (
	"github.com/vnchk1/CalculatorAPI/configs"
	"log/slog"
	"os"
)

func HandlerInit() *slog.JSONHandler {
	cfg := configs.LoadConfig()
	logLevel := configs.ConvertLogLevel(cfg.LoggerLevel)
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	})
	return logHandler
}

func Logger() *slog.Logger {
	logger := slog.New(HandlerInit())
	return logger
}
