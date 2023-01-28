package controllers

import (
	"strconv"

	"github.com/DiarCode/todo-go-api/src/config/database"
	"github.com/DiarCode/todo-go-api/src/dto"
	"github.com/DiarCode/todo-go-api/src/helpers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllTodos(c *fiber.Ctx) error {
	user_param := c.Query("user")

	if user_param == "" {
		todos := []Todo{}
		database.DB.Find(&todos)
		return helpers.SendSuccessJSON(c, todos)
	}

	userId, err := strconv.Atoi(user_param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Provide user id in params Format",
		})
	}

	todos := []Todo{}
	database.DB.Where("user_id = ?", userId).Find(&todos)

	return helpers.SendSuccessJSON(c, todos)
}

func GetTodoById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID Format",
		})
	}

	todo := Todo{}
	query := Todo{ID: id}
	err = database.DB.First(&todo, &query).Error

	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Todo not found",
		})
	}

	return helpers.SendSuccessJSON(c, todo)
}

func GetTodoByCategory(c *fiber.Ctx) error {
	json := new(dto.TodoByCategoryDto)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	todos := []Todo{}
	query := Todo{CategoryId: json.CategoryId, UserId: json.UserId}
	database.DB.Find(&todos, &query)

	return helpers.SendSuccessJSON(c, todos)
}

func CreateTodo(c *fiber.Ctx) error {
	json := new(dto.CreateTodoDto)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	newTodo := Todo{
		UserId:     json.UserId,
		Title:      json.Title,
		Priority:   json.Priority,
		CategoryId: json.CategoryId,
	}

	err := database.DB.Create(&newTodo).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return helpers.SendSuccessJSON(c, newTodo)
}

func DeleteTodoById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID format",
		})
	}

	foundTodo := Todo{}
	query := Todo{
		ID: id,
	}

	err = database.DB.First(&foundTodo, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Todo not found",
		})
	}

	database.DB.Delete(&foundTodo)
	return helpers.SendSuccessJSON(c, nil)
}

func CompleteTodoById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID format",
		})
	}

	foundTodo := Todo{}
	query := Todo{
		ID: id,
	}

	err = database.DB.First(&foundTodo, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Todo not found",
		})
	}

	database.DB.Delete(&foundTodo)
	return helpers.SendSuccessJSON(c, nil)
}
