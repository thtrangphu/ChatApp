package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mekanican/chat-backend/internal/database"
	"github.com/mekanican/chat-backend/internal/model"
)

func GetAllRoom(c *fiber.Ctx) error {
	var rooms []model.Room
	database.DB.Find(&rooms)
	return c.JSON(rooms)
}

func GetRoom(c *fiber.Ctx) error {
	id := c.Params("id")
	var roomInfo model.Room
	database.DB.Find(&roomInfo, id)
	return c.JSON(roomInfo)
}

func CreateRoom(c *fiber.Ctx) error {
	var roomInfo model.Room
	if err := c.BodyParser(&roomInfo); err != nil {
		return c.Status(500).Send([]byte("Failed"))
	}

	database.DB.Create(&roomInfo)
	return c.JSON(roomInfo)
}

func DeleteRoom(c *fiber.Ctx) error {
	id := c.Params("id")
	var roomInfo model.Room
	database.DB.First(&roomInfo, id)
	database.DB.Delete(&roomInfo)
	return c.JSON(roomInfo)
}
