package middleware

import (
	"os"
	"strings"

	"github.com/DiarCode/todo-go-api/src/helpers"
	"github.com/DiarCode/todo-go-api/src/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := c.Next(); err != nil {
			return err
		}

		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return helpers.SendMessageWithStatus(c, "Autharization header not provided", 404)
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return helpers.SendMessageWithStatus(c, "Invalid autharization header", 404)
		}

		token := headerParts[1]

		jwtKey := os.Getenv("JWT_KEY")
		claims := &models.Claims{}

		tkn, err := jwt.ParseWithClaims(token, claims,
			func(t *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return helpers.SendMessageWithStatus(c, "Unautharized", 401)
			}

			return helpers.SendMessageWithStatus(c, "Bad request", 404)
		}

		if !tkn.Valid {
			return helpers.SendMessageWithStatus(c, "Unautharized", 401)
		}

		return helpers.SendMessageWithStatus(c, "Unautharized", 401)
	}
}
