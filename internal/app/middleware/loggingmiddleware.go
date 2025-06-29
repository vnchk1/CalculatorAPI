package middleware

import (
	"github.com/labstack/echo/v4"
	"log/slog"
)

func LoggingMiddleware(logger *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			logger.Info("middleware is working",
				"method", c.Request().Method)
			if err != nil {
				logger.Error("request is failed", "error", err)
			}
			return nil
		}
	}
}
