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
	credentials := new(dto.LoginDto)
	if err := c.BodyParser(credentials); err != nil {
		return utils.SendMessageWithStatus(c, "Invalid JSON", 400)
	}

	user := User{}
	query := User{Email: credentials.Email}
	err := database.DB.First(&user, &query).Error

	if err == gorm.ErrRecordNotFound {
		return utils.SendMessageWithStatus(c, "User not found", 404)
	}

	if !utils.ComparePasswords(user.Password, credentials.Password) {
		return utils.SendMessageWithStatus(c, "Passwords does not match", 404)
	}

	tokenString, err := generateToken(user)

	if err != nil {
		return utils.SendMessageWithStatus(c, "Auth error (token creation)", 500)
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
	json := new(dto.SignupDto)
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

	err = database.DB.Create(&newUser).Error
	if err != nil {
		return utils.SendMessageWithStatus(c, err.Error(), 400)
	}

	tokenString, err := generateToken(newUser)

	if err != nil {
		return utils.SendMessageWithStatus(c, "Auth error (token creation)", 500)
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
