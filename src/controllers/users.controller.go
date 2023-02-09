package controllers

import (
	"strconv"

	"github.com/DiarCode/todo-go-api/src/database"
	"github.com/DiarCode/todo-go-api/src/dto"
	"github.com/DiarCode/todo-go-api/src/models"
	"github.com/DiarCode/todo-go-api/src/utils"
	"github.com/badoux/checkmail"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllUsers(c *fiber.Ctx) error {
	users := []User{}
	database.DB.Model(&models.User{}).Preload("Todos").Find(&users)

	return utils.SendSuccessJSON(c, users)
}

func GetUserById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return utils.SendMessageWithStatus(c, "Invalid ID Format", 400)
	}

	user := User{}
	query := User{ID: id}
	err = database.DB.First(&user, &query).Error

	if err == gorm.ErrRecordNotFound {
		return utils.SendMessageWithStatus(c, "User not found", 404)
	}

	return utils.SendSuccessJSON(c, user)
}

func CreateUser(c *fiber.Ctx) error {
	json := new(dto.CreateUserDto)
	if err := c.BodyParser(json); err != nil {
		return utils.SendMessageWithStatus(c, "Invalid JSON", 400)
	}

	password := utils.HashPassword([]byte(json.Password))
	err := checkmail.ValidateFormat(json.Email)
	if err != nil {
		return utils.SendMessageWithStatus(c, "Invalid Email Address", 400)
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
		return utils.SendMessageWithStatus(c, "User already exists", 400)
	}

	database.DB.Create(&newUser)

	return utils.SendSuccessJSON(c, newUser)
}

func DeleteUserById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return utils.SendMessageWithStatus(c, "Invalid ID format", 400)
	}

	user := User{}
	query := User{
		ID: id,
	}

	err = database.DB.First(&user, &query).Error
	if err == gorm.ErrRecordNotFound {
		return utils.SendMessageWithStatus(c, "User not found", 400)
	}

	database.DB.Model(&user).Association("Todos").Delete()
	database.DB.Delete(&user)
	return utils.SendSuccessJSON(c, nil)
}
