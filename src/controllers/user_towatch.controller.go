package controllers

import (
	"github.com/DiarCode/todo-go-api/src/config/database"
	"github.com/DiarCode/todo-go-api/src/dto"
	"github.com/DiarCode/todo-go-api/src/helpers"
	"github.com/DiarCode/todo-go-api/src/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllTowathesByCategory(c *fiber.Ctx) error {
	json := new(dto.UserTowatchDto)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	towatchCategory := TowatchCategory{}
	towatchCategoryQuery := TowatchCategory{ID: json.CategoryID}
	err := database.DB.First(&towatchCategory, &towatchCategoryQuery).Error

	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Towatch category not found",
		})
	}

	userTowatch := UserTowatch{}
	query := UserTowatch{TowatchCategory: models.TowatchCategory(towatchCategory), UserID: json.UserID}
	err = database.DB.First(&userTowatch, &query).Error

	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "User Towatch not found",
		})
	}

	return helpers.SendSuccessJSON(c, userTowatch.Towatches)
}

func AssignTowatchToCategory(c *fiber.Ctx) error {
	json := new(dto.AssignTowatchToCategoryDto)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	towatchCategory := TowatchCategory{}
	towatchCategoryQuery := TowatchCategory{ID: json.CategoryID}
	err := database.DB.First(&towatchCategory, &towatchCategoryQuery).Error

	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Towatch category not found",
		})
	}

	towatch := Towatch{}
	towatchQuery := Towatch{ID: json.TowatchID}
	err = database.DB.First(&towatch, &towatchQuery).Error

	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Towatch not found",
		})
	}	

	userTowatch := UserTowatch{}
	query := UserTowatch{TowatchCategory: models.TowatchCategory(towatchCategory), UserID: json.UserID}
	err = database.DB.First(&userTowatch, &query).Error

	if err == gorm.ErrRecordNotFound {
		newUserTowatch := UserTowatch{
			UserID:          json.UserID,
			TowatchCategory: models.TowatchCategory(towatchCategory),
		}

		err := database.DB.Create(&newUserTowatch).Error
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		userTowatch = newUserTowatch
	}

	towatches := userTowatch.Towatches 
	for _, t := range towatches {
		if t.ID == userTowatch.ID {
			return c.SendStatus(fiber.StatusBadRequest)
		}
	}

	towatches = append(towatches, models.Towatch(towatch))
	userTowatch.Towatches = towatches
	err = database.DB.Save(&userTowatch).Error
	if err != nil {
		return helpers.SendMessageWithStatus(c, "Could not add towatch", 500)
	}

	return helpers.SendSuccessJSON(c, towatches)

}
