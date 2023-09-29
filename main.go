package main

import (
	"os"
	"todo/configs"
	"todo/routes"
	_ "todo/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// @title Swagger Todo Application
// @version 1.0
// @description This is a todo application.
// @termsOfService http://swagger.io/terms/

// @contact.name Achmad Kumail Arif
// @contact.url https://github.com/ackuarif
// @contact.email ackuarif@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host todoapp.osc-fr1.scalingo.io
// @BasePath /

func main(){
	//loadEnv()
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoutes(e)
	e.Start(":"+os.Getenv("PORT"))
}

func loadEnv(){
	err := godotenv.Load()
  	if err != nil {
    	panic("Error loading .env file")
  	}
}


