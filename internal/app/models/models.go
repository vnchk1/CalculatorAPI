package models

// @Description NumbersRequest представляет входной JSON с массивом чисел и токеном
type NumbersRequest struct {
	Numbers []int  `json:"numbers" example:"1,2,4"`
	Token   string `json:"token" example:"82a51b41-b7c7-4405-859e-7936xxxxxxxx"`
}

// @Description SumResponse представляет исходящий JSON с целым числом
type SumResponse struct {
	Sum int `json:"sum" example:"7"`
}

// @Description ErrorResponse представляет исходящий JSON со строкой ошибки ("Invalid request body" или "Empty request body")
type ErrorResponse struct {
	ErrorRequest string `json:"error" example:"Invalid request body"`
}

// @Description MultiplyResponse представляет исходящий JSON с целым числом
type MultiplyResponse struct {
	Multiply int `json:"multiply" example:"8"`
}
