package routes

import (
	controllers "todo/controllers"

	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/todo", controllers.GetAllTodosController)
	e.GET("/todo/:id", controllers.GetTodoByIdController)
	e.POST("/todo", controllers.AddTodoController)
	e.PUT("/todo/:id", controllers.UpdateTodoController)
	e.DELETE("/todo/:id", controllers.DeleteTodoController)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}