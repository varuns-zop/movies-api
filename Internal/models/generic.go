package models

type GenericResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   string `json:"data"`
}
