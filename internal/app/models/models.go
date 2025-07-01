package models

type NumbersRequest struct {
	Numbers []int  `json:"numbers"`
	Token   string `json:"token"`
}

type SumResponse struct {
	Sum int `json:"sum"`
}

type MultiplyResponse struct {
	Multiply int `json:"multiply"`
}
