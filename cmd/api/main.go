package main

import (
	"github.com/labstack/gommon/log"
	"github.com/vnchk1/CalculatorAPI/internal/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// middleware.Logger логирует каждый запрос, а отдельные логгеры для ошибок
func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/sum", handler.SumHandler)
	e.Logger.Fatal(e.Start(":8080"))

}
