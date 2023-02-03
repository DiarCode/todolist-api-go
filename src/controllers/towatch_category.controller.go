package controllers

import (
	"strconv"

	"github.com/DiarCode/todo-go-api/src/database"
	"github.com/DiarCode/todo-go-api/src/dto"
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

	categories := []TodoCategory{}
	query := TodoCategory{UserId: userId}
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
