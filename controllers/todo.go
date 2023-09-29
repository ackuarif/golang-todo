package controllers

import (
	"net/http"
	"todo/configs"
	models "todo/models"

	"github.com/labstack/echo/v4"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// GetAllTodo godoc
// @Summary      Menampilkan semua list todo
// @Description  get all todo
// @Tags         todo
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Response{data=[]models.Todo}
// @Failure      500  {object}  models.ResponseFailed
// @Router       /todo [get]
func GetAllTodosController(c echo.Context) error{
	var todos []models.Todo
	result := configs.DB.Order("id DESC").Find(&todos)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseFailed{
			Status: false,
			Message: "Failed get data from database",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status: true,
		Message: "Success",
		Data: todos,
	})
}

// GetTodoById godoc
// @Summary      Menampilkan todo berdasarkan ID
// @Description  get todo by id
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param id path int true "Todo ID"
// @Success      200  {object}  models.Response{data=models.Todo}
// @Failure      500  {object}  models.ResponseFailed
// @Router       /todo/{id} [get]
func GetTodoByIdController(c echo.Context) error{
	todo := new(models.Todo)
	result := configs.DB.Where("id = ?", c.Param("id")).First(&todo)
	
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseFailed{
			Status: false,
			Message: "Failed get data from database",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status: true,
		Message: "Success",
		Data: todo,
	})
}

// AddTodo godoc
// @Summary      Menambah todo
// @Description  add todo
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param request body models.Todo true "raw body"
// @Success      200  {object}  models.Response{data=models.Todo}
// @Failure      500  {object}  models.ResponseFailed
// @Router       /todo [post]
func AddTodoController(c echo.Context) error {
	addTodo := new(models.Todo)
	c.Bind(&addTodo)

	validate = validator.New()
    if validationErr := validate.Struct(addTodo); validationErr != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseFailed{
			Status: false,
			Message: "Incomplete parameter.",
		})
	}

	result := configs.DB.Create(&addTodo)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseFailed{
			Status: false,
			Message: "Failed insert data to database",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status: true,
		Message: "Success insert data to database",
		Data: addTodo,
	})
}

// UpdateTodo godoc
// @Summary      Merubah todo berdasarkan ID
// @Description  update todo by id
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param id path int true "Todo ID"
// @Param request body models.Todo true "raw body"
// @Success      200  {object}  models.Response{data=models.Todo}
// @Failure      500  {object}  models.ResponseFailed
// @Router       /todo/{id} [put]
func UpdateTodoController(c echo.Context) error {
	updateTodo := new(models.Todo)
	getTodo := new(models.Todo)
	c.Bind(&updateTodo)

	result := configs.DB.Where("id = ?", c.Param("id")).First(&getTodo)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseFailed{
			Status: false,
			Message: "Failed update data to database",
		})
	}

	configs.DB.Where("id = ?", c.Param("id")).Updates(&updateTodo)

	return c.JSON(http.StatusOK, models.Response{
		Status: true,
		Message: "Success update data to database",
		Data: updateTodo,
	})
}

// DeleteTodo godoc
// @Summary      Merubah todo berdasarkan ID
// @Description  delete todo by id
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param id path int true "Todo ID"
// @Success      200  {object}  models.Response{data=models.Todo}
// @Failure      500  {object}  models.ResponseFailed
// @Router       /todo/{id} [delete]
func DeleteTodoController(c echo.Context) error {
	var deleteTodo models.Todo

	configs.DB.Where("id = ?", c.Param("id")).First(&deleteTodo)
	result := configs.DB.Delete(deleteTodo)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseFailed{
			Status: false,
			Message: "Data not found.",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status: true,
		Message: "Success delete data from database",
		Data: deleteTodo,
	})
}