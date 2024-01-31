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
		return utils.ResponseError(c, "Invalid body "+err.Error())
	}

	database.DB.Create(&userInfo)
	return c.JSON(userInfo)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ResponseError(c, "Invalid ID")
	}

	// var userInfo model.User

	var info map[string]interface{}

	if err := c.BodyParser(&info); err != nil {
		return utils.ResponseError(c, "Invalid body "+err.Error())
	}

	// info["ID"] = uint(id)
	if err := database.DB.Model(&model.User{}).Where("id = ?", uint(id)).Updates(&info).Error; err != nil {
		return utils.ResponseError(c, "Invalid body "+err.Error())
	}
	return c.JSON(info)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ResponseError(c, "Invalid ID")
	}
	var userInfo model.User
	if err := database.DB.First(&userInfo, uint(id)).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.ResponseError(c, "Invalid User")
	}
	database.DB.Delete(&userInfo)
	return c.JSON(userInfo)
}

// --------------------------------------------------------------
func CheckUserInDB(email string) bool {

	if err := database.DB.First(&model.User{}, model.User{Email: email}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
