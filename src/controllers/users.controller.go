package controllers

import (
	"strconv"

	"github.com/DiarCode/todo-go-api/src/config/database"
	"github.com/DiarCode/todo-go-api/src/dto"
	"github.com/DiarCode/todo-go-api/src/helpers"
	"github.com/DiarCode/todo-go-api/src/models"
	"github.com/badoux/checkmail"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllUsers(c *fiber.Ctx) error {
	users := []User{}
	database.DB.Model(&models.User{}).Preload("Todos").Find(&users)

	return helpers.SendSuccessJSON(c, users)
}

func GetUserById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID Format",
		})
	}

	user := User{}
	query := User{ID: id}
	err = database.DB.First(&user, &query).Error

	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "User not found",
		})
	}

	return helpers.SendSuccessJSON(c, user)
}

func CreateUser(c *fiber.Ctx) error {
	json := new(dto.CreateUserDto)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	password := helpers.HashPassword([]byte(json.Password))
	err := checkmail.ValidateFormat(json.Email)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Email Address",
		})
	}

	newUser := User{
		Password: password,
		Email:    json.Email,
	}

	found := User{}
	query := User{Email: json.Email}
	err = database.DB.First(&found, &query).Error
	if err != gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "User already exists",
		})
	}

	database.DB.Create(&newUser)

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
		"data":    newUser,
	})
}

func DeleteUserById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID format",
		})
	}

	user := User{}
	query := User{
		ID: id,
	}

	err = database.DB.First(&user, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "User not found",
		})
	}

	database.DB.Model(&user).Association("Todos").Delete()
	database.DB.Delete(&user)
	return helpers.SendSuccessJSON(c, nil)
}
