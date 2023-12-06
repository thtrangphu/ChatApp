package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/mekanican/chat-backend/docs"
)

//	@Summary	Get hello in response
//	@ID			get-hello
//	@Produce	plain
//	@Success	200	{string}	hello
//	@Router		/api/v1/hello [get]
func getHello(c*fiber.Ctx) error {
  return c.SendString("Hello")
}

//	@Summary	Get hello in response with ID
//	@ID			get-hello-id
//	@Param		id	path	string	true	"random ID"
//	@Produce	plain
//	@Success	200	{string}	hello
//	@Router		/api/v1/hello/{id} [get]
func getHelloId(c *fiber.Ctx) error {
  msg := fmt.Sprintf("Hello ID: %s", c.Params("id"))
  return c.SendString(msg)
}

//	@title			Go + Fiber Chat app API
//	@version		0.1
//	@description	This is a sample server. You can visit the GitHub repository at https://github.com/thtrangphu/ChatApp
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

//	@host						localhost:8000
//	@BasePath					/api/v1
//	@query.collection.format	multi

func main() {
  app := fiber.New()

  // Route for swagger
  app.Get("/swagger/*", swagger.HandlerDefault)

  v1 := app.Group("/api/v1")
  {
    hello := v1.Group("/hello")
    {
      hello.Get("", getHello)
      hello.Get(":id", getHelloId)
    }
  }

  // Listen on PORT 8000
  app.Listen(":8000")
}
