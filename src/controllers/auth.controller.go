package controllers

import (
	"os"
	"time"

	"github.com/DiarCode/todo-go-api/src/config/database"
	"github.com/DiarCode/todo-go-api/src/dto"
	"github.com/DiarCode/todo-go-api/src/helpers"
	"github.com/badoux/checkmail"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Token struct {
	Token string `json:"token"`
}

type Claims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func Login(c *fiber.Ctx) error {
	credentials := new(dto.LoginDto)
	if err := c.BodyParser(credentials); err != nil {
		return helpers.SendMessageWithStatus(c, "Invalid JSON", 400)
	}

	user := User{}
	query := User{Email: credentials.Email}
	err := database.DB.First(&user, &query).Error

	if err == gorm.ErrRecordNotFound {
		return helpers.SendMessageWithStatus(c, "User not found", 404)
	}

	passwordsMatch := helpers.ComparePasswords(user.Password, []byte(credentials.Password))

	if !passwordsMatch {
		return helpers.SendMessageWithStatus(c, "Passwords does not match", 404)
	}

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
	tokenString, err := token.SignedString(jwtKey)

	sendToken := &Token{
		Token: tokenString,
	}

	if err != nil {
		return helpers.SendMessageWithStatus(c, "Auth error (token creation)", 500)
	}

	return helpers.SendSuccessJSON(c, sendToken)
}

func Signup(c *fiber.Ctx) error {
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
