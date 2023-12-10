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

	if isAuth {
		return c.SendString("Authenticated User!")
	}
	return c.SendString("Unauthenticated!")
}

func IndexRoute(router fiber.Router) {
	router.Use(utils.AuthMiddleware)
	router.Get("/", SampleIndex)
}
