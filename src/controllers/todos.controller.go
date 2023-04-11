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
		return fiber.NewError(fiber.StatusBadRequest, "Provide proper user id")
	}

	todos := []Todo{}
	database.DB.Where("user_id = ?", userId).Find(&todos)

	return utils.SendSuccessJSON(c, todos)
}

func GetTodosByCategory(c *fiber.Ctx) error {
	category_param := c.Params("id")
	user_param := c.Query("user")

	if user_param == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Provide user id in params")
	}

	userId, err := strconv.Atoi(user_param)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Provide proper user id")

	}

	categoryId, err := strconv.Atoi(category_param)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Provide proper category id")
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
		return fiber.NewError(fiber.StatusBadRequest, "Invalid id")
	}

	todo := Todo{}
	query := Todo{ID: id}
	err = database.DB.First(&todo, &query).Error

	if err == gorm.ErrRecordNotFound {
		return fiber.NewError(fiber.StatusNotFound, "Todo not found")
	}

	return utils.SendSuccessJSON(c, todo)
}

func CreateTodo(c *fiber.Ctx) error {
	var json dto.CreateTodoDto
	err := c.BodyParser(&json)
	if err != nil || (json == dto.CreateTodoDto{}) {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body")
	}

	newTodo := Todo{
		UserId:     json.UserId,
		Title:      json.Title,
		Priority:   json.Priority,
		CategoryId: json.CategoryId,
	}

	err = database.DB.Create(&newTodo).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create todo")
	}

	return utils.SendSuccessJSON(c, newTodo)
}

func DeleteTodoById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Provide proper id")
	}

	foundTodo := Todo{}
	query := Todo{
		ID: id,
	}

	err = database.DB.First(&foundTodo, &query).Error
	if err == gorm.ErrRecordNotFound {
		return fiber.NewError(fiber.StatusNotFound, "Todo not found")
	}

	database.DB.Delete(&foundTodo)
	return utils.SendSuccessJSON(c, nil)
}

func CompleteTodoById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Provide proper id")
	}

	foundTodo := Todo{}
	query := Todo{
		ID: id,
	}

	err = database.DB.First(&foundTodo, &query).Error
	if err == gorm.ErrRecordNotFound {
		return fiber.NewError(fiber.StatusNotFound, "Todo not found")
	}

	database.DB.Delete(&foundTodo)
	return utils.SendSuccessJSON(c, nil)
}
