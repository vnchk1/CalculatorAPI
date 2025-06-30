package middleware

import (
	"github.com/labstack/echo/v4"
	"log/slog"
)

func LoggingMiddleware(logger *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				logger.Error("request is failed", "error", err)
			}
			logger.Info("request ended", "status", c.Response().Status,
				"method", c.Request().Method)
			return nil
		}
	}
}
