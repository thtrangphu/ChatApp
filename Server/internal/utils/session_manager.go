package utils

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var serverStore = session.New(session.Config{
	Expiration: 12 * time.Hour, // Session expires after 12 hours
})

func SetKV(ctx *fiber.Ctx, key string, value interface{}, expiryHour time.Duration) error {
	sess, err := serverStore.Get(ctx)
	if err != nil {
		return err
	}

	sess.Set(key, value)
	sess.SetExpiry(expiryHour * time.Hour)
	if err = sess.Save(); err != nil {
		return err
	}
	return nil
}

func GetKV[T uint | string](ctx *fiber.Ctx, key string) (T, error) {
	sess, err := serverStore.Get(ctx)
	var nilResult T
	if err != nil {
		return nilResult, err
	}

	id, ok := sess.Get(key).(T)
	if !ok {
		// TODO: Replace with meaningful errors
		return nilResult, errors.ErrUnsupported
	}
	return id, nil
}

func DestroySession(ctx *fiber.Ctx) error {
	sess, err := serverStore.Get(ctx)
	if err != nil {
		return err
	}

	return sess.Destroy()
}

func AuthMiddleware(ctx *fiber.Ctx) error {
	email, err := GetKV[string](ctx, "Email")
	if err != nil {
		return ResponseError(ctx, "User unauthenticated!")
	}
	ctx.Locals("Email", email)
	return ctx.Next() //  TODO: use database to validate
}

func CheckUserMiddleware(ctx *fiber.Ctx) error {
	id, err := GetKV[uint](ctx, "UserID")

	if err != nil {
		return ResponseError(ctx, "User not registered")
	}
	ctx.Locals("UserID", id)
	return ctx.Next()
}
