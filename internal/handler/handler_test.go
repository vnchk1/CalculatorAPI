package handler

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/vnchk1/CalculatorAPI/internal/store"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSumHandler(t *testing.T) {
	e := echo.New()

	logger := slog.Default()
	storage := store.NewSafeMap()
	h := NewHandler(logger, storage)

	tests := []struct {
		name        string
		requestBody map[string]interface{}
		wantStatus  int
		wantBody    map[string]interface{}
	}{
		{
			name: "valid json",
			requestBody: map[string]interface{}{
				"numbers": []int{1, 2, 3},
				"token":   "token123",
			},
			wantStatus: http.StatusOK,
			wantBody: map[string]interface{}{
				"sum": float64(6),
			},
		},
		{
			name:        "invalid json",
			requestBody: nil,
			wantStatus:  http.StatusBadRequest,
			wantBody: map[string]interface{}{
				"error": "Invalid request body",
			},
		},
		{
			name: "empty array",
			requestBody: map[string]interface{}{
				"numbers": []int{},
				"token":   "token123",
			},
			wantStatus: http.StatusBadRequest,
			wantBody: map[string]interface{}{
				"error": "Empty request body",
			},
		},
		{
			name: "invalid token",
			requestBody: map[string]interface{}{
				"numbers": []int{1, 2, 3},
				"token":   "",
			},
			wantStatus: http.StatusBadRequest,
			wantBody: map[string]interface{}{
				"error": "Invalid request body",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reqBody []byte
			var err error
			//создаем json для тела запроса
			if tt.requestBody != nil {
				reqBody, err = json.Marshal(tt.requestBody)
				assert.NoError(t, err)
			} else {
				reqBody = []byte("Invalid json")
			}
			//создаем запрос и устанавливаем заголовок
			req := httptest.NewRequest(http.MethodPost, "/sum", bytes.NewBuffer(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			//io.Writer для ответа
			rec := httptest.NewRecorder()
			//констекст echo
			c := e.NewContext(req, rec)
			//работа хендлера
			err = h.SumHandler(c)
			//проверка на ошибку и сравнение
			assert.NoError(t, err)
			assert.Equal(t, tt.wantStatus, rec.Code)
			//объект для ответа
			var resp map[string]interface{}
			//анмаршаллим тело ответа
			err = json.Unmarshal(rec.Body.Bytes(), &resp)
			//проверка на ошибку и сравнение
			assert.NoError(t, err)
			assert.Equal(t, tt.wantBody, resp)
		})
	}
}

func TestMultiplyHandler(t *testing.T) {
	e := echo.New()

	logger := slog.Default()
	storage := store.NewSafeMap()
	h := NewHandler(logger, storage)

	tests := []struct {
		name        string
		requestBody map[string]interface{}
		wantStatus  int
		wantBody    map[string]interface{}
	}{
		{
			name: "valid json",
			requestBody: map[string]interface{}{
				"numbers": []int{1, 2, 4},
				"token":   "token123",
			},
			wantStatus: http.StatusOK,
			wantBody: map[string]interface{}{
				"multiply": float64(8),
			},
		},
		{
			name:        "invalid json",
			requestBody: nil,
			wantStatus:  http.StatusBadRequest,
			wantBody: map[string]interface{}{
				"error": "Invalid request body",
			},
		},
		{
			name: "empty array",
			requestBody: map[string]interface{}{
				"numbers": []int{},
				"token":   "token123",
			},
			wantStatus: http.StatusBadRequest,
			wantBody: map[string]interface{}{
				"error": "Empty request body",
			},
		},
		{
			name: "invalid token",
			requestBody: map[string]interface{}{
				"numbers": []int{1, 2, 3},
				"token":   "",
			},
			wantStatus: http.StatusBadRequest,
			wantBody: map[string]interface{}{
				"error": "Invalid request body",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reqBody []byte
			var err error
			if tt.requestBody != nil {
				reqBody, err = json.Marshal(tt.requestBody)
			} else {
				reqBody = []byte("Invalid json")
			}

			req := httptest.NewRequest(http.MethodPost, "/multiply", bytes.NewBuffer(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			err = h.MultiplyHandler(c)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantStatus, rec.Code)

			var resp map[string]interface{}
			err = json.Unmarshal(rec.Body.Bytes(), &resp)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantBody, resp)
		})
	}
}
