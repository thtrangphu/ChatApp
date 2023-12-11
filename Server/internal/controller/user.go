package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mekanican/chat-backend/internal/database"
	"github.com/mekanican/chat-backend/internal/model"
)

func GetAllUser(c *fiber.Ctx) error {
	var users []model.User
	database.DB.Find(&users)
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var userInfo model.User
	database.DB.Find(&userInfo, id)
	return c.JSON(userInfo)
}

func CreateUser(c *fiber.Ctx) error {
	var userInfo model.User
	if err := c.BodyParser(&userInfo); err != nil {
		return c.Status(500).Send([]byte("Failed"))
	}

	database.DB.Create(&userInfo)
	return c.JSON(userInfo)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var userInfo model.User
	database.DB.First(&userInfo, id)
	database.DB.Delete(&userInfo)
	return c.JSON(userInfo)
}
