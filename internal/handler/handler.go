package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/vnchk1/CalculatorAPI/internal/app/models"
	"net/http"
)

func SumHandler(c echo.Context) error {
	var req models.NumbersRequest

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
	return c.JSON(http.StatusOK, models.SumResponse{Sum: sum})
}
