package models

type User struct {
	Id int 		`json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type ResponseAuth struct {
	Id int 		`json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}