package controllers

import (
	"log"
	"strconv"

	"github.com/DiarCode/todo-go-api/src/database"
	"github.com/DiarCode/todo-go-api/src/dto"
	"github.com/DiarCode/todo-go-api/src/models"
	"github.com/DiarCode/todo-go-api/src/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllTowatch(c *fiber.Ctx) error {
	towatches := []Towatch{}
	database.DB.Model(&models.Towatch{}).Find(&towatches)

	return utils.SendSuccessJSON(c, towatches)
}

func GetTowatchesByCategory(c *fiber.Ctx) error {
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

	towatchCategory := TowatchCategory{}
	query := TowatchCategory{ID: categoryId, UserId: userId}
	database.DB.Preload("Towatches").First(&towatchCategory, query)

	log.Println(towatchCategory.Towatches)
	return utils.SendSuccessJSON(c, towatchCategory.Towatches)
}

func GetTowatchById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID Format",
		})
	}

	towatch := Towatch{}
	query := Towatch{ID: id}
	err = database.DB.First(&towatch, &query).Error

	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Todo not found",
		})
	}

	return utils.SendSuccessJSON(c, towatch)
}

func CreateTowatch(c *fiber.Ctx) error {
	json := new(dto.CreateTowatchDto)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	newTowatch := Towatch{
		Title:      json.Title,
		StartDate:  json.StartDate,
		FinishDate: json.FinishDate,
		Episodes:   json.Episodes,
		Rating:     json.Rating,
		Studio:     json.Studio,
		Image:      json.Image,
	}

	err := database.DB.Create(&newTowatch).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return utils.SendSuccessJSON(c, newTowatch)
}

func DeleteTowatchById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID format",
		})
	}

	foundTowatch := Towatch{}
	query := Towatch{
		ID: id,
	}

	err = database.DB.First(&foundTowatch, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Towatch not found",
		})
	}

	database.DB.Delete(&foundTowatch)
	return utils.SendSuccessJSON(c, nil)
}
