package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type NumbersRequest struct {
	Numbers []int `json:"numbers"`
}

type SumResponse struct {
	Sum int `json:"sum"`
}

func SumHandler(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		"method", r.Method,
	)
	if r.Method != http.MethodPost {
		logger.Error("Method not allowed")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var req NumbersRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Error with JSON decoding", "error", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var sum int
	for _, n := range req.Numbers {
		sum += n
	}

	resp := SumResponse{sum}
	w.Header().Set(http.CanonicalHeaderKey("content-type"), "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		logger.Error("Error with JSON encoding", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	logger.Info(fmt.Sprintf("Sum calculated: %v", sum))
}
func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	slog.SetDefault(slog.New(handler))
	http.HandleFunc("/sum", SumHandler)
	fmt.Println("Listening on port 8080")
	slog.Error("Server stopped", "error", http.ListenAndServe(":8080", nil))
}
