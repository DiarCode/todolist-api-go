package controllers

import (
	"os"
	"time"

	"github.com/DiarCode/todo-go-api/src/database"
	"github.com/DiarCode/todo-go-api/src/dto"
	"github.com/DiarCode/todo-go-api/src/utils"
	"github.com/badoux/checkmail"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	var credentials dto.LoginDto
	err := c.BodyParser(&credentials)

	if err != nil || (credentials == dto.LoginDto{}) {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body")
	}

	user := User{}
	query := User{Email: credentials.Email}
	err = database.DB.First(&user, &query).Error

	if err == gorm.ErrRecordNotFound {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	if !utils.ComparePasswords(user.Password, credentials.Password) {
		return fiber.NewError(fiber.StatusBadRequest, "Passwords does not match")

	}
	tokenString, err := generateToken(user)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Auth error (token creation)")
	}

	response := &TokenResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: tokenString,
	}

	return utils.SendSuccessJSON(c, response)
}

func Signup(c *fiber.Ctx) error {
	var json dto.SignupDto
	err := c.BodyParser(&json)

	if err != nil || (json == dto.SignupDto{}) {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Body")
	}

	password := utils.HashPassword([]byte(json.Password))
	err = checkmail.ValidateFormat(json.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Email Address")
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
		return fiber.NewError(fiber.StatusBadRequest, "User already exists")
	}

	err = database.DB.Create(&newUser).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create new user")
	}

	tokenString, err := generateToken(newUser)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Auth error (token creation)")
	}

	response := &TokenResponse{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
		Token: tokenString,
	}

	return utils.SendSuccessJSON(c, response)
}

func generateToken(user User) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)

	claims := &Claims{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	jwtKey := os.Getenv("JWT_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))

	return tokenString, err
}
