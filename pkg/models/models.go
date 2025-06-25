package models

type NumbersRequest struct {
	Numbers []int `json:"numbers"`
}

type SumResponse struct {
	Sum int `json:"sum"`
}
