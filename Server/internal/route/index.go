package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mekanican/chat-backend/internal/utils"
)

func SampleIndex(c *fiber.Ctx) error {
	isAuth, ok := c.Locals("authenticated").(bool)
	if !ok {
		return c.SendStatus(502)
	}
	isInDB, ok := c.Locals("isInDB").(bool)
	if !ok {
		return c.SendStatus(502)
	}

	if isAuth {
		if isInDB {
			return c.SendString("Authenticated User!")
		} else {
			return c.SendString("First time user!")
		}
	}
	return c.SendString("Unauthenticated!")
}

func IndexRoute(router fiber.Router) {
	router.Use(utils.AuthMiddleware)
	router.Get("/", SampleIndex)
}
