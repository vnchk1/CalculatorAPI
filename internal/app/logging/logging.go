package logging

import (
	"github.com/vnchk1/CalculatorAPI/configs"
	"log/slog"
	"os"
)

func NewLogger(cfg *configs.Config) *slog.Logger {
	logLevel := ConvertLogLevel(cfg.LoggerLevel)
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	})
	return slog.New(logHandler)
}

func ConvertLogLevel(lvlStr string) slog.Level {
	switch lvlStr {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
