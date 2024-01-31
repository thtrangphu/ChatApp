package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mekanican/chat-backend/internal/utils"
)

func SampleIndex(c *fiber.Ctx) error {
	return c.SendString("Hello " + c.Locals("Email").(string))
}

func IndexRoute(router fiber.Router) {
	// router.Use(utils.AuthMiddleware)
	router.Get("/", utils.AuthMiddleware, SampleIndex)
}
