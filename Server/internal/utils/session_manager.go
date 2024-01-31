package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var serverStore = session.New(session.Config{
	Expiration: 12 * time.Hour, // Session expires after 12 hours
})

func SetKV(ctx *fiber.Ctx, key string, value string, expiryHour time.Duration) error {
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

func GetKV(ctx *fiber.Ctx, key string) (string, error) {
	sess, err := serverStore.Get(ctx)
	if err != nil {
		return "", err
	}

	id, ok := sess.Get(key).(string)
	if !ok {
		return "", nil
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
	id, err := GetKV(ctx, "Email")
	if err != nil {
		ctx.Locals("authenticated", false)
		return ctx.Next()
	}

	isInDB, err := GetKV(ctx, "IsInDB")
	if err != nil {
		ctx.Locals("IsInDB", false)
		return ctx.Next()
	}

	ctx.Locals("isInDB", isInDB != "")
	ctx.Locals("authenticated", id != "")
	return ctx.Next() //  TODO: use database to validate
}
