package models

//ResponseSuccess representa o response caso retorne sucesso
type Response struct {
	Mensage string      `json:"meta"`
	Status  string      `json:"status"`
	Log     interface{} `json:"records"`
}
