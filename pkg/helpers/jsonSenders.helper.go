package helpers

import "github.com/gofiber/fiber/v2"

func SendSuccessJSON(c *fiber.Ctx, data any) error {
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
		"data":    data,
	})
}
