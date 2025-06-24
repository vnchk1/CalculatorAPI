package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
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
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	var sum int
	for _, num := range req.Numbers {
		sum += num
	}

	return c.JSON(http.StatusOK, SumResponse{Sum: sum})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/sum", SumHandler)
	e.Logger.Fatal(e.Start(":8080"))

}
