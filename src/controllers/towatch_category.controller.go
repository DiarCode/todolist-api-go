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
		return fiber.NewError(fiber.StatusBadRequest, "Provide user id in params")
	}

	userId, err := strconv.Atoi(user_param)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Provide proper user id in params")
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
		return fiber.NewError(fiber.StatusBadRequest, "Provide proper category id in params")
	}

	category := TowatchCategory{}
	query := TowatchCategory{ID: id}
	err = database.DB.First(&category, &query).Error

	if err == gorm.ErrRecordNotFound {
		return fiber.NewError(fiber.StatusNotFound, "Towatch category not found")
	}

	return utils.SendSuccessJSON(c, category)
}

func CreateTowatchCategory(c *fiber.Ctx) error {
	var json dto.CreateTowatchCategoryDto
	err := c.BodyParser(&json)
	if err != nil || (json == dto.CreateTowatchCategoryDto{}) {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body")
	}

	newCategory := TowatchCategory{
		Value:  json.Value,
		Color:  json.Color,
		UserId: json.UserId,
	}

	err = database.DB.Create(&newCategory).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create a towatch category")
	}

	return utils.SendSuccessJSON(c, newCategory)
}

func DeleteTowatchCategoryById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Provide proper towatch id in params")
	}

	foundCategory := TowatchCategory{}
	query := TowatchCategory{
		ID: id,
	}

	err = database.DB.First(&foundCategory, &query).Error
	if err == gorm.ErrRecordNotFound {
		return fiber.NewError(fiber.StatusNotFound, "Towatch category not found")
	}

	database.DB.Delete(&foundCategory)
	return utils.SendSuccessJSON(c, nil)
}

func AddTowatchToCategory(c *fiber.Ctx) error {
	var json dto.AddTowatchToCategoryDto
	err := c.BodyParser(json)
	if err != nil || (json == dto.AddTowatchToCategoryDto{}) {
		return fiber.NewError(fiber.StatusBadRequest, "Invlalid body")

	}

	towatchCategory := TowatchCategory{}
	category_query := TowatchCategory{ID: json.TowatchCategoryId, UserId: json.UserId}
	err = database.DB.Preload("Towatches").First(&towatchCategory, &category_query).Error

	if err == gorm.ErrRecordNotFound {
		return fiber.NewError(fiber.StatusNotFound, "Towatch category not found")
	}

	towatch := Towatch{}
	towatch_query := Towatch{ID: json.TowatchId}
	err = database.DB.First(&towatch, &towatch_query).Error

	if err == gorm.ErrRecordNotFound {
		return fiber.NewError(fiber.StatusNotFound, "Towatch not found")

	}

	towatches := towatchCategory.Towatches
	for _, t := range towatches {
		if t.ID == towatch.ID {
			return fiber.NewError(fiber.StatusBadRequest, "Towatch already added")
		}
	}

	towatchCategory.Towatches = append(towatchCategory.Towatches, models.Towatch(towatch))
	err = database.DB.Save(towatchCategory).Error

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Couldn't add towatch to category")
	}

	return utils.SendSuccessJSON(c, towatchCategory)
}

func RemoveTowatchFromCategory(c *fiber.Ctx) error {
	var json dto.AddTowatchToCategoryDto
	err := c.BodyParser(json)
	if err != nil || (json == dto.AddTowatchToCategoryDto{}) {
		return fiber.NewError(fiber.StatusBadRequest, "Invlalid body")
	}

	towatchCategory := TowatchCategory{}
	category_query := TowatchCategory{ID: json.TowatchCategoryId, UserId: json.UserId}
	err = database.DB.Preload("Towatches").First(&towatchCategory, &category_query).Error

	if err == gorm.ErrRecordNotFound {
		return fiber.NewError(fiber.StatusNotFound, "Towatch category not found")
	}

	towatch := Towatch{}
	towatch_query := Towatch{ID: json.TowatchId}
	err = database.DB.First(&towatch, &towatch_query).Error

	if err == gorm.ErrRecordNotFound {
		return fiber.NewError(fiber.StatusNotFound, "Towatch not found")
	}

	towatches := towatchCategory.Towatches
	filteredTowatches := []models.Towatch{}

	for _, t := range towatches {
		if t.ID != towatch.ID {
			filteredTowatches = append(filteredTowatches, t)
		}
	}

	database.DB.Model(&towatchCategory).Association("Towatches").Clear()
	towatchCategory.Towatches = filteredTowatches
	err = database.DB.Save(&towatchCategory).Error

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Couldn't remove towatch to category")
	}

	return utils.SendSuccessJSON(c, towatchCategory)
}
