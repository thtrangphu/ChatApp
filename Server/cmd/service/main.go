package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New()

  v1 := app.Group("/api/v1")
  {
    hello := v1.Group("/hello")
    {
      hello.Get("", func(c *fiber.Ctx) error {
        return c.SendString("Hello")
      })

      hello.Get(":id", func(c *fiber.Ctx) error {
        msg := fmt.Sprintf("Hello ID: %s", c.Params("id"))
        return c.SendString(msg)
      })
    }
  }

  // Listen on PORT 8000
  app.Listen(":8000")
}
