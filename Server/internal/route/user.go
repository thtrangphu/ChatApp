package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mekanican/chat-backend/internal/controller"
)

func UserRoute(router fiber.Router) {
	// router.Use(utils.AuthMiddleware)
	router.Get("/", controller.GetAllUser)
	router.Get("/:id", controller.GetUser)
	router.Post("/", controller.CreateUser)
	router.Delete("/:id", controller.DeleteUser)
}
