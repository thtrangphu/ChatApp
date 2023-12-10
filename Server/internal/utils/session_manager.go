package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var serverStore = session.New(session.Config{
	Expiration: 12 * time.Hour, // Session expires after 12 hours
})

func SetID(ctx *fiber.Ctx, id string) error {
	sess, err := serverStore.Get(ctx)
	if err != nil {
		return err
	}

	sess.Set("ID", id)
	sess.SetExpiry(4 * time.Hour)
	if err = sess.Save(); err != nil {
		return err
	}
	return nil
}

func GetId(ctx *fiber.Ctx) (string, error) {
	sess, err := serverStore.Get(ctx)
	if err != nil {
		return "", err
	}

	id, ok := sess.Get("ID").(string)
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
	id, err := GetId(ctx)
	if err != nil {
		ctx.Locals("authenticated", false)
		return ctx.Next()
	}

	ctx.Locals("authenticated", id != "")
	return ctx.Next() //  TODO: use database to validate
}
