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
	if r.Method != http.MethodPost {
		log.Printf("Method %v not allowed", r.Method)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var req NumbersRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error with JSON decoding: %v", err)
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
		log.Printf("Error with JSON encoding: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
func main() {
	http.HandleFunc("/sum", SumHandler)
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
