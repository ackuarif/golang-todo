package models

type Response struct {
	Status bool 		`json:"status,omitempty" validate:"required" example:true`
	Message string 		`json:"message,omitempty" validate:"required" example:"Success"`
	Data interface{} 	`json:"data,omitempty"`
}

type ResponseFailed struct {
	Status bool 	`json:"status,omitempty" validate:"required" example:false`
	Message interface{} `json:"message,omitempty"`
}