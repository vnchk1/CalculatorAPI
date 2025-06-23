package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type NumbersRequest struct {
	Numbers []int `json:"numbers"`
}

type SumResponse struct {
	Sum int `json:"sum"`
}

func SumHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var req NumbersRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	var sum int
	for _, n := range req.Numbers {
		sum += n
	}

	resp := SumResponse{sum}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

}
func main() {
	http.HandleFunc("/sum", SumHandler)
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
