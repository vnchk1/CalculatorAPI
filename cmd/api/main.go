package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vnchk1/CalculatorAPI/configs"
	"github.com/vnchk1/CalculatorAPI/internal/app/logging"
	"github.com/vnchk1/CalculatorAPI/internal/app/middleware"
	"github.com/vnchk1/CalculatorAPI/internal/handler"
	"log/slog"
)

// middleware.Logger логирует каждый запрос, а отдельные логгеры для ошибок
func main() {
	e := echo.New()
	cfg, err := configs.LoadConfig()
	if err != nil {
		slog.Error("Error loading config", "err", err)
	}
	//logLvlStr := cfg.Logger.Level
	//logLevel := configs.ConvertLogLevel(logLvlStr)

	port := cfg.Server.Port
	//logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	//	Level: logLevel,
	//})
	//logger := slog.New(logHandler)
	//e.Logger.SetLevel(configs.ConvertLogLevel(logLvlStr))
	logger := logging.Logger()
	e.Use(middleware.LoggingMiddleware(logger))
	//e.Use(middleware.Recover())
	e.POST("/sum", handler.SumHandler)
	e.Logger.Fatal(e.Start(":" + port))
}
