package controllers

import (
	"strconv"

	"github.com/DiarCode/todo-go-api/src/config/database"
	"github.com/DiarCode/todo-go-api/src/dto"
	"github.com/DiarCode/todo-go-api/src/helpers"
	"github.com/DiarCode/todo-go-api/src/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllTodoCategories(c *fiber.Ctx) error {
	categories := []TodoCategory{}
	database.DB.Model(&models.TodoCategory{}).Find(&categories)

	return helpers.SendSuccessJSON(c, categories)
}

func GetTodoCategoryById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID Format",
		})
	}

	category := TodoCategory{}
	query := TodoCategory{ID: id}
	err = database.DB.First(&category, &query).Error

	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Todo category not found",
		})
	}

	return helpers.SendSuccessJSON(c, category)
}

func CreateTodoCategory(c *fiber.Ctx) error {
	json := new(dto.CreateTodoCategoryDto)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	newCategory := TodoCategory{
		Value: json.Value,
		Color: json.Color,
	}

	err := database.DB.Create(&newCategory).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return helpers.SendSuccessJSON(c, newCategory)
}

func DeleteTodoCategoryById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID format",
		})
	}

	foundCategory := TodoCategory{}
	query := TodoCategory{
		ID: id,
	}

	err = database.DB.First(&foundCategory, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Todo category not found",
		})
	}

	database.DB.Delete(&foundCategory)
	return helpers.SendSuccessJSON(c, nil)
}
