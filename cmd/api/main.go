package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vnchk1/CalculatorAPI/configs"
	"github.com/vnchk1/CalculatorAPI/internal/app/logging"
	"github.com/vnchk1/CalculatorAPI/internal/app/middleware"
	"github.com/vnchk1/CalculatorAPI/internal/handler"
	"github.com/vnchk1/CalculatorAPI/internal/store"
)

func main() {
	e := echo.New()

	cfg := configs.LoadConfig()
	logger := logging.NewLogger(cfg)
	storage := store.GetStorage()

	e.Use(middleware.LoggingMiddleware(logger))
	//e.Use(middleware.Recover())

	h := handler.NewHandler(logger, storage)

	e.POST("/sum", h.SumHandler)
	e.POST("/multiply", h.MultiplyHandler)

	e.Logger.Fatal(e.Start(":" + cfg.ServerPort))
}
