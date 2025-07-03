package handler

import (
	"log/slog"
	"net/http"

	"github.com/vnchk1/CalculatorAPI/internal/app/models"
	"github.com/vnchk1/CalculatorAPI/internal/app/service"
	"github.com/vnchk1/CalculatorAPI/internal/store"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	logger  *slog.Logger
	storage *store.SafeMap
}

func NewHandler(logger *slog.Logger, storage *store.SafeMap) *Handler {
	return &Handler{logger: logger, storage: storage}
}

// sum godoc
// @Summary Сложение чисел
// @Description Складывает все числа из запроса и сохраняется результат по токену
// @Tags sum
// @Accept json
// @Produce json
// @Param data body models.NumbersRequest true "Массив чисел и токен пользователя"
// @Success 200 {object} models.SumResponse "Успешный результат с суммой"
// @Failure 400 {object} models.ErrorResponse "Ошибка: "Invalid request body" или "Empty request body" "
// @Router /sum [post]
func (h *Handler) SumHandler(c echo.Context) error {
	var req models.NumbersRequest
	//id := uuid.New().String()

	if err := c.Bind(&req); err != nil || req.Token == "" {
		h.logger.Error("error with parsing JSON", "error", err)
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{ErrorRequest: "Invalid request body"})
	}

	sum, err := service.Sum(req.Numbers)
	id := req.Token

	if err != nil {
		h.logger.Error("calculating sum error", "error", err)
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{ErrorRequest: "Empty request body"})
	}
	h.storage.MapSet(id, sum)
	return c.JSON(http.StatusOK, models.SumResponse{Sum: sum})
}

// multiply godoc
// @Summary Произведение чисел
// @Description Умножает все числа из массива запроса
// @Tags multiply
// @Accept json
// @Produce json
// @Param data body models.NumbersRequest true "Массив чисел и токен пользователя"
// @Success 200 {object} models.MultiplyResponse "Успешный результат с произведением"
// @Failure 400 {object} models.ErrorResponse "Ошибка: "Invalid request body" или "Empty request body" "
// @Router /multiply [post]
func (h *Handler) MultiplyHandler(c echo.Context) error {
	var req models.NumbersRequest
	//id := uuid.New().String()

	if err := c.Bind(&req); err != nil || req.Token == "" {
		h.logger.Error("error with parsing JSON", "error", err)
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{ErrorRequest: "Invalid request body"})
	}

	multiply, err := service.Multiply(req.Numbers)
	id := req.Token

	if err != nil {
		h.logger.Error("calculating multiply error", "error", err)
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{ErrorRequest: "Empty request body"})
	}
	h.storage.MapSet(id, multiply)
	return c.JSON(http.StatusOK, models.MultiplyResponse{Multiply: multiply})
}
