package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/vnchk1/CalculatorAPI/internal/app/models"
	"github.com/vnchk1/CalculatorAPI/internal/app/service"
	"net/http"
)

func SumHandler(c echo.Context) error {
	var req models.NumbersRequest

	if err := c.Bind(&req); err != nil {
		c.Logger().Error("Error with parsing JSON ", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}
	sum, err := service.Sum(req.Numbers)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, models.SumResponse{Sum: sum})
}
