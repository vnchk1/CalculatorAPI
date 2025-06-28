package handler

import (
	"net/http"

	"github.com/vnchk1/CalculatorAPI/internal/app/models"
	"github.com/vnchk1/CalculatorAPI/internal/app/service"
	"github.com/vnchk1/CalculatorAPI/internal/store"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func SumHandler(c echo.Context) error {
	var req models.NumbersRequest
	id := uuid.New()

	if err := c.Bind(&req); err != nil {
		c.Logger().Error("Error with parsing JSON ", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}
	sum, err := service.Sum(req.Numbers)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	storage := store.GetStorage()
	storage.MapSet(id, sum)
	//fmt.Println(storage.MapGet(id))
	//fmt.Println(storage.MapGetAll())
	return c.JSON(http.StatusOK, models.SumResponse{Sum: sum})
}
