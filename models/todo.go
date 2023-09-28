package models

type Todo struct {
	Id int 		`json:"id,omitempty"`
	Todo string `json:"todo" validate:"required" example:"This is my todo"`
	Date string `json:"date" validate:"required" example:"2023-01-03"`
}