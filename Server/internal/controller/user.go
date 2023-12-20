package controller

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mekanican/chat-backend/internal/database"
	"github.com/mekanican/chat-backend/internal/model"
	"github.com/mekanican/chat-backend/internal/utils"
	"gorm.io/gorm"
)

func GetAllUser(c *fiber.Ctx) error {
	var users []model.User
	database.DB.Find(&users)
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ResponseError(c, "Invalid ID")
	}

	var userInfo model.User
	if err := database.DB.First(&userInfo, uint(id)).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.ResponseError(c, "Invalid User")
	}
	// database.DB.Find(&userInfo, id)
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
