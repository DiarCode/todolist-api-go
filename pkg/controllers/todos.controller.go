package controllers

import (
	"strconv"

	"github.com/DiarCode/todo-go-api/pkg/dto"
	"github.com/DiarCode/todo-go-api/pkg/helpers"
	"github.com/DiarCode/todo-go-api/pkg/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllTodos(c *fiber.Ctx) error {
	todos := []Todo{}
	db.Model(&models.Todo{}).Find(&todos)

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
	err = db.First(&todo, &query).Error

	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Todo not found",
		})
	}

	return helpers.SendSuccessJSON(c, todo)
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
		UserRefer:   json.UserId,
		Title:       json.Title,
		Description: json.Description,
		Completed:   false,
	}

	err := db.Create(&newTodo).Error
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

	err = db.First(&foundTodo, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Todo not found",
		})
	}

	db.Delete(&foundTodo)
	return helpers.SendSuccessJSON(c, nil)
}
