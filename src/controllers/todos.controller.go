package controllers

import (
	"strconv"

	"github.com/DiarCode/todo-go-api/src/database"
	"github.com/DiarCode/todo-go-api/src/dto"
	"github.com/DiarCode/todo-go-api/src/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllTodos(c *fiber.Ctx) error {
	user_param := c.Query("user")

	if user_param == "" {
		todos := []Todo{}
		database.DB.Find(&todos)
		return utils.SendSuccessJSON(c, todos)
	}

	userId, err := strconv.Atoi(user_param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Provide user id in proper params",
		})
	}

	todos := []Todo{}
	database.DB.Where("user_id = ?", userId).Find(&todos)

	return utils.SendSuccessJSON(c, todos)
}

func GetTodosByCategory(c *fiber.Ctx) error {
	category_param := c.Params("id")
	user_param := c.Query("user")

	if user_param == "" {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Provide user id in params",
		})
	}

	userId, err := strconv.Atoi(user_param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Provide user id in proper params",
		})
	}

	categoryId, err := strconv.Atoi(category_param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Provide category id in proper format",
		})
	}

	todos := []Todo{}
	query := Todo{CategoryId: categoryId, UserId: userId}
	database.DB.Find(&todos, query)

	return utils.SendSuccessJSON(c, todos)
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

	return utils.SendSuccessJSON(c, todo)
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

	return utils.SendSuccessJSON(c, newTodo)
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
	return utils.SendSuccessJSON(c, nil)
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
	return utils.SendSuccessJSON(c, nil)
}
