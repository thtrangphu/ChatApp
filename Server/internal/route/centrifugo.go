package route

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mekanican/chat-backend/internal/utils"
)

func AuthNToken(ctx *fiber.Ctx) error {
	val := ctx.Locals("UserID").(uint)
	tok, err := utils.GenerateRoomJWT(strconv.Itoa(int(val)), "", 1)
	if err != nil {
		return utils.ResponseError(ctx, err.Error())
	}
	return ctx.Status(200).JSON(map[string]string{"token": tok})
}

func AuthZToken(ctx *fiber.Ctx) error {
	val := ctx.Locals("UserID").(uint)
	tok, err := utils.GenerateRoomJWT(strconv.Itoa(int(val)), "chat", 1)
	if err != nil {
		return utils.ResponseError(ctx, err.Error())
	}
	return ctx.Status(200).JSON(map[string]string{"token": tok})
}

func CentrifugoRoute(router fiber.Router) {
	router.Get("/token-n", utils.AuthMiddleware, utils.CheckUserMiddleware, AuthNToken)
	router.Get("/token-z", utils.AuthMiddleware, utils.CheckUserMiddleware, AuthZToken)
}
