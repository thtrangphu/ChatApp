package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mekanican/chat-backend/apiproto"
	"github.com/mekanican/chat-backend/internal/utils"
)

// TODO: Use db to get correct room
func SampleChat(c *fiber.Ctx) error {
	var info map[string]interface{}
	if err := c.BodyParser(&info); err != nil {
		return utils.ResponseError(c, "Invalid body "+err.Error())
	}

	mess, ok := info["message"].(string)
	if !ok {
		return utils.ResponseError(c, "Invalid message")
	}

	// TODO: handle DB
	err := apiproto.PublishData("sample-chat", apiproto.MessageData{
		Sender:  c.Locals("Email").(string),
		Message: mess,
	})

	if err != nil {
		return utils.ResponseError(c, "Cannot publish "+err.Error())
	}
	return nil

	// return c.SendFile("./samplechat/index.html")
	// return c.SendString("Hello " + c.Locals("Email").(string))
}

// TODO: Use db to get correct room
func GetRoom(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{"room": "sample-room"})
}

func ChatRoute(router fiber.Router) {
	// router.Use(utils.AuthMiddleware)
	router.Get("/getroom", utils.AuthMiddleware, utils.CheckUserMiddleware, GetRoom)
	router.Post("/", utils.AuthMiddleware, utils.CheckUserMiddleware, SampleChat)
}
