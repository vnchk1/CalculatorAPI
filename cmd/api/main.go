package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/vnchk1/CalculatorAPI/configs"
	_ "github.com/vnchk1/CalculatorAPI/docs"
	"github.com/vnchk1/CalculatorAPI/internal/app/logging"
	"github.com/vnchk1/CalculatorAPI/internal/app/middleware"
	"github.com/vnchk1/CalculatorAPI/internal/handler"
	"github.com/vnchk1/CalculatorAPI/internal/store"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title CalculatorAPI
// @version 1.0
// @description Это API для вычисления суммы и произведения

// @host localhost:8080
// @BasePath /

// @tag.name sum
// @tag.description Вычисление суммы

// @tag.name multiply
// @tag.description Вычисление произведения

func main() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	cfg := configs.LoadConfig()
	logger := logging.NewLogger(cfg)
	storage := store.GetStorage()

	e.Use(middleware.LoggingMiddleware(logger))
	//e.Use(middleware.Recover())

	h := handler.NewHandler(logger, storage)

	e.POST("/sum", h.SumHandler)
	e.POST("/multiply", h.MultiplyHandler)

	go func() {
		err := e.Start(":" + cfg.ServerPort)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("error starting server", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mapCopy := storage.MapGetAll()
	for key, value := range mapCopy {
		fmt.Printf("%v : %v \n", key, value)
	}

	if err := e.Shutdown(ctx); err != nil {
		logger.Error("server forced to shutdown", "error", err)
	} else {
		logger.Info("server gracefully shutdown")
	}
}
