package handler

import (
	"github.com/vnchk1/CalculatorAPI/internal/app/logging"
	"net/http"

	"github.com/vnchk1/CalculatorAPI/internal/app/models"
	"github.com/vnchk1/CalculatorAPI/internal/app/service"
	"github.com/vnchk1/CalculatorAPI/internal/store"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func SumHandler(c echo.Context) error {
	logger := logging.Logger()
	var req models.NumbersRequest
	id := uuid.New()

	if err := c.Bind(&req); err != nil {
		logger.Error("Error with parsing JSON", "error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}
	sum, err := service.Sum(req.Numbers)
	if err != nil {
		logger.Error("Calculating sum problem", "error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Empty request body"})
	}
	storage := store.GetStorage()
	storage.MapSet(id, sum)
	//fmt.Println(storage.MapGet(id))
	//GetAllMaps := storage.MapGetAll())
	return c.JSON(http.StatusOK, models.SumResponse{Sum: sum})
}
