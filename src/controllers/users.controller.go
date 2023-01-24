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
		return helpers.SendMessageWithStatus(c, "Invalid ID Format", 400)
	}

	user := User{}
	query := User{ID: id}
	err = database.DB.First(&user, &query).Error

	if err == gorm.ErrRecordNotFound {
		return helpers.SendMessageWithStatus(c, "User not found", 404)
	}

	return helpers.SendSuccessJSON(c, user)
}

func CreateUser(c *fiber.Ctx) error {
	json := new(dto.CreateUserDto)
	if err := c.BodyParser(json); err != nil {
		return helpers.SendMessageWithStatus(c, "Invalid JSON", 400)
	}

	password := helpers.HashPassword([]byte(json.Password))
	err := checkmail.ValidateFormat(json.Email)
	if err != nil {
		return helpers.SendMessageWithStatus(c, "Invalid Email Address", 400)
	}

	newUser := User{
		Password: password,
		Email:    json.Email,
		Name:     json.Name,
	}

	found := User{}
	query := User{Email: json.Email}
	err = database.DB.First(&found, &query).Error
	if err != gorm.ErrRecordNotFound {
		return helpers.SendMessageWithStatus(c, "User already exists", 400)
	}

	database.DB.Create(&newUser)

	return helpers.SendSuccessJSON(c, newUser)
}

func DeleteUserById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return helpers.SendMessageWithStatus(c, "Invalid ID format", 400)
	}

	user := User{}
	query := User{
		ID: id,
	}

	err = database.DB.First(&user, &query).Error
	if err == gorm.ErrRecordNotFound {
		return helpers.SendMessageWithStatus(c, "User not found", 400)
	}

	database.DB.Model(&user).Association("Todos").Delete()
	database.DB.Delete(&user)
	return helpers.SendSuccessJSON(c, nil)
}
