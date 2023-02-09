package utils

import "github.com/gofiber/fiber/v2"

func SendSuccessJSON(c *fiber.Ctx, data any) error {
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
		"data":    data,
	})
}

func SendMessageWithStatus(c *fiber.Ctx, message string, status int) error {
	return c.JSON(fiber.Map{
		"code":    status,
		"message": message,
	})
}
