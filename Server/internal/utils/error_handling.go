package utils

import "github.com/gofiber/fiber/v2"

func ResponseError(c *fiber.Ctx, message string) error {
	return c.Status(500).JSON(map[string]string{"error": message})
}
