package models

type Response struct {
	Status bool 		`json:"status" validate:"required" example:true`
	Message string 		`json:"message" validate:"required" example:"Success"`
	Data interface{} 	`json:"data"`
}

type ResponseFailed struct {
	Status bool 	`json:"status" validate:"required" example:false`
	Message interface{} `json:"message"`
}