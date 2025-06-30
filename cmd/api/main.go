package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vnchk1/CalculatorAPI/configs"
	"github.com/vnchk1/CalculatorAPI/internal/app/logging"
	"github.com/vnchk1/CalculatorAPI/internal/app/middleware"
	"github.com/vnchk1/CalculatorAPI/internal/handler"
)

func main() {
	e := echo.New()
	logger := logging.Logger()
	cfg := configs.LoadConfig()
	port := cfg.ServerPort
	e.Use(middleware.LoggingMiddleware(logger))
	//e.Use(middleware.Recover())
	e.POST("/sum", handler.SumHandler)
	e.POST("/multiply", handler.MultiplyHandler)
	e.Logger.Fatal(e.Start(":" + port))
}
