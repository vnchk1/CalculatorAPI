package main

import (
	"github.com/labstack/gommon/log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type NumbersRequest struct {
	Numbers []int `json:"numbers"`
}

type SumResponse struct {
	Sum int `json:"sum"`
}

func SumHandler(c echo.Context) error {
	var req NumbersRequest

	if err := c.Bind(&req); err != nil {
		c.Logger().Error("Error with parsing JSON ", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}
	if len(req.Numbers) < 1 {
		c.Logger().Error("Empty list of numbers")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Empty list of numbers"})
	}

	var sum int
	for _, num := range req.Numbers {
		sum += num
	}
	return c.JSON(http.StatusOK, SumResponse{Sum: sum})
}

// middleware.Logger логирует каждый запрос, а отдельные логгеры для ошибок и успешного выполнения запроса
func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/sum", SumHandler)
	e.Logger.Fatal(e.Start(":8080"))

}
