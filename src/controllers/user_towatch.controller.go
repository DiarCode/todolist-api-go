package controllers

import (
	"github.com/DiarCode/todo-go-api/src/config/database"
	"github.com/DiarCode/todo-go-api/src/dto"
	"github.com/DiarCode/todo-go-api/src/helpers"
	"github.com/DiarCode/todo-go-api/src/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllTowatchesByCategory(c *fiber.Ctx) error {
	json := new(dto.UserTowatchDto)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	userTowatch := UserTowatch{}
	query := UserTowatch{TowatchCategoryID: json.CategoryID, UserID: json.UserID}
	err := database.DB.Model(&models.UserTowatch{}).Preload("Towatches").Preload("TowatchCategory").First(&userTowatch, &query).Error

	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Towatch category not found",
		})
	}

	return helpers.SendSuccessJSON(c, userTowatch)
}

func AssignTowatchToCategory(c *fiber.Ctx) error {
	json := new(dto.AssignTowatchToCategoryDto)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	towatch := Towatch{}
	towatchQuery := Towatch{ID: json.TowatchID}
	err := database.DB.First(&towatch, &towatchQuery).Error

	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Towatch not found",
		})
	}

	userTowatch := UserTowatch{}
	query := UserTowatch{TowatchCategoryID: json.CategoryID, UserID: json.UserID}
	err = database.DB.Model(&models.UserTowatch{}).Preload("Towatches").First(&userTowatch, &query).Error

	if err == gorm.ErrRecordNotFound {
		newUserTowatch := UserTowatch{
			UserID:            json.UserID,
			TowatchCategoryID: json.CategoryID,
		}

		err := database.DB.Create(&newUserTowatch).Error
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		userTowatch = newUserTowatch
	}

	towatches := userTowatch.Towatches
	for _, t := range towatches {
		if t.ID == towatch.ID {
			return helpers.SendMessageWithStatus(c, "Towtach already added", 400)
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

func RemoveTowatchFromCategory(c *fiber.Ctx) error {
	json := new(dto.RemoveTowatchFromCategoryDto)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	towatch := Towatch{}
	towatchQuery := Towatch{ID: json.TowatchID}
	err := database.DB.First(&towatch, &towatchQuery).Error

	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Towatch not found",
		})
	}

	userTowatch := UserTowatch{}
	query := UserTowatch{TowatchCategoryID: json.CategoryID, UserID: json.UserID}
	err = database.DB.Model(&models.UserTowatch{}).Preload("Towatches").Preload("").First(&userTowatch, &query).Error

	if err == gorm.ErrRecordNotFound {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(fiber.Map{
				"code":    404,
				"message": "User Towatch not found",
			})
		}
	}

	towatches := userTowatch.Towatches
	filteredTowatches := []models.Towatch{}

	for _, t := range towatches {
		if t.ID != towatch.ID {
			filteredTowatches = append(filteredTowatches, t)
		}
	}

	// // log.Println("filtered", filteredTowatches)
	userTowatch.Towatches = filteredTowatches
	// // // log.Println("userTowatch", userTowatch.Towatches)
	// // err = database.DB.Save(&userTowatch).Error

	// sql_query := fmt.Sprintf("DELETE FROM user_towatch_cards WHERE user_towatch_id = %v AND towatch_id = %v", userTowatch.ID, towatch.ID)
	// log.Println(sql_query)
	// err = database.DB.Raw(sql_query).Error\

	err = database.DB.Updates(&userTowatch).Error
	if err != nil {
		return helpers.SendMessageWithStatus(c, err.Error(), 500)
	}

	return helpers.SendSuccessJSON(c, userTowatch)
}
