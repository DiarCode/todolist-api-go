package controllers

import (
	"strconv"

	"github.com/DiarCode/todo-go-api/src/database"
	"github.com/DiarCode/todo-go-api/src/dto"
	"github.com/DiarCode/todo-go-api/src/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllTodoCategories(c *fiber.Ctx) error {
	user_param := c.Query("user")
	if user_param == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Empty user id")
	}

	userId, err := strconv.Atoi(user_param)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Provide proper user id")
	}

	categories := []TodoCategory{}
	query := TodoCategory{UserId: userId}
	database.DB.Find(&categories, query)

	return utils.SendSuccessJSON(c, categories)
}

func GetTodoCategoryById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Provide proper category id")
	}
	category := TodoCategory{}
	query := TodoCategory{ID: id}
	err = database.DB.First(&category, &query).Error

	if err == gorm.ErrRecordNotFound {
		return fiber.NewError(fiber.StatusNotFound, "Todo category not found")
	}

	return utils.SendSuccessJSON(c, category)
}

func CreateTodoCategory(c *fiber.Ctx) error {
	var json dto.CreateTodoCategoryDto
	err := c.BodyParser(&json)
	if err != nil || (json == dto.CreateTodoCategoryDto{}) {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body")
	}

	newCategory := TodoCategory{
		Value:  json.Value,
		Color:  json.Color,
		UserId: json.UserId,
	}

	err = database.DB.Create(&newCategory).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create a todo category")
	}

	return utils.SendSuccessJSON(c, newCategory)
}

func DeleteTodoCategoryById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Provide proper category id")
	}

	foundCategory := TodoCategory{}
	query := TodoCategory{
		ID: id,
	}

	err = database.DB.First(&foundCategory, &query).Error
	if err == gorm.ErrRecordNotFound {
		return fiber.NewError(fiber.StatusNotFound, "Todo category not found")
	}

	todos := []Todo{}
	query_todos := Todo{CategoryId: foundCategory.ID}

	database.DB.Find(&todos, &query_todos)

	database.DB.Delete(&todos)
	database.DB.Delete(&foundCategory)
	return utils.SendSuccessJSON(c, nil)
}
