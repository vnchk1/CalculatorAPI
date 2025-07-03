package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/vnchk1/CalculatorAPI/configs"
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

func main() {
	e := echo.New()
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
