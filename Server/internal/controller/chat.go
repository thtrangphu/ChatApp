package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mekanican/chat-backend/internal/database"
	"github.com/mekanican/chat-backend/internal/model"
)

func GetAllChat(c *fiber.Ctx) error {
	var chats []model.Chat
	database.DB.Find(&chats)
	return c.JSON(chats)
}

func GetChat(c *fiber.Ctx) error {
	id := c.Params("id")
	var chatInfo model.Chat
	database.DB.Find(&chatInfo, id)
	return c.JSON(chatInfo)
}

func CreateChat(c *fiber.Ctx) error {
	var chatInfo model.Chat
	if err := c.BodyParser(&chatInfo); err != nil {
		return c.Status(500).Send([]byte("Failed"))
	}

	database.DB.Create(&chatInfo)
	return c.JSON(chatInfo)
}

func DeleteChat(c *fiber.Ctx) error {
	id := c.Params("id")
	var chatInfo model.Chat
	database.DB.First(&chatInfo, id)
	database.DB.Delete(&chatInfo)
	return c.JSON(chatInfo)
}
