package controllers

import (
	"strconv"

	"github.com/DiarCode/todo-go-api/src/database"
	"github.com/DiarCode/todo-go-api/src/dto"
	"github.com/DiarCode/todo-go-api/src/models"
	"github.com/DiarCode/todo-go-api/src/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllTowatchCategories(c *fiber.Ctx) error {
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
			"message": "Provide user id in params",
		})
	}

	categories := []TowatchCategory{}
	query := TowatchCategory{UserId: userId}
	database.DB.Find(&categories, query)

	return utils.SendSuccessJSON(c, categories)
}

func GetTowatchCategoryById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID Format",
		})
	}

	category := TowatchCategory{}
	query := TowatchCategory{ID: id}
	err = database.DB.First(&category, &query).Error

	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Towatch category not found",
		})
	}

	return utils.SendSuccessJSON(c, category)
}

func CreateTowatchCategory(c *fiber.Ctx) error {
	json := new(dto.CreateTowatchCategoryDto)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	newCategory := TowatchCategory{
		Value:  json.Value,
		Color:  json.Color,
		UserId: json.UserId,
	}

	err := database.DB.Create(&newCategory).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return utils.SendSuccessJSON(c, newCategory)
}

func DeleteTowatchCategoryById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID format",
		})
	}

	foundCategory := TowatchCategory{}
	query := TowatchCategory{
		ID: id,
	}

	err = database.DB.First(&foundCategory, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Towatch category not found",
		})
	}

	database.DB.Delete(&foundCategory)
	return utils.SendSuccessJSON(c, nil)
}

func AddTowatchToCategory(c *fiber.Ctx) error {
	json := new(dto.AddTowatchToCategoryDto)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	towatchCategory := TowatchCategory{}
	category_query := TowatchCategory{ID: json.TowatchCategoryId, UserId: json.UserId}
	err := database.DB.Preload("Towatches").First(&towatchCategory, &category_query).Error

	if err == gorm.ErrRecordNotFound {
		return utils.SendMessageWithStatus(c, "Towatch category not found", 404)

	}

	towatch := Towatch{}
	towatch_query := Towatch{ID: json.TowatchId}
	err = database.DB.First(&towatch, &towatch_query).Error

	if err == gorm.ErrRecordNotFound {
		return utils.SendMessageWithStatus(c, "Towatch not found", 404)

	}

	towatches := towatchCategory.Towatches
	for _, t := range towatches {
		if t.ID == towatch.ID {
			return utils.SendMessageWithStatus(c, "Towatch already added", 400)
		}
	}

	towatchCategory.Towatches = append(towatchCategory.Towatches, models.Towatch(towatch))
	err = database.DB.Save(towatchCategory).Error

	if err != nil {
		return utils.SendMessageWithStatus(c, "Couldn't add towatch to category", 404)
	}

	return utils.SendSuccessJSON(c, towatchCategory)
}

func RemoveTowatchFromCategory(c *fiber.Ctx) error {
	json := new(dto.AddTowatchToCategoryDto)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	towatchCategory := TowatchCategory{}
	category_query := TowatchCategory{ID: json.TowatchCategoryId, UserId: json.UserId}
	err := database.DB.Preload("Towatches").First(&towatchCategory, &category_query).Error

	if err == gorm.ErrRecordNotFound {
		return utils.SendMessageWithStatus(c, "Towatch category not found", 404)

	}

	towatch := Towatch{}
	towatch_query := Towatch{ID: json.TowatchId}
	err = database.DB.First(&towatch, &towatch_query).Error

	if err == gorm.ErrRecordNotFound {
		return utils.SendMessageWithStatus(c, "Towatch not found", 404)

	}

	towatches := towatchCategory.Towatches
	filteredTowatches := []models.Towatch{}

	for _, t := range towatches {
		if t.ID != towatch.ID {
			filteredTowatches = append(filteredTowatches, t)
		}
	}

	towatchCategory.Towatches = filteredTowatches
	err = database.DB.Save(towatchCategory).Error

	if err != nil {
		return utils.SendMessageWithStatus(c, "Couldn't remove towatch to category", 404)
	}

	return utils.SendSuccessJSON(c, towatchCategory)
}