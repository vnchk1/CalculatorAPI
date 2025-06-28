package main

import (
	"github.com/vnchk1/CalculatorAPI/configs"
	"github.com/vnchk1/CalculatorAPI/internal/handler"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// middleware.Logger логирует каждый запрос, а отдельные логгеры для ошибок
func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	logLvlStr := cfg.Logger.Level
	port := cfg.Server.Port
	e := echo.New()
	e.Logger.SetLevel(configs.ConvertLogLevel(logLvlStr))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/sum", handler.SumHandler)
	e.Logger.Fatal(e.Start(":" + port))
}
