package models

type ExecuteRequest struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}
