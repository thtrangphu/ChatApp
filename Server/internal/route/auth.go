package route

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mekanican/chat-backend/internal/controller"
	"github.com/mekanican/chat-backend/internal/oauth"
	"github.com/mekanican/chat-backend/internal/utils"
)

// @Summary Redirect client to google signin
// @ID			redirect-google
// @Produce	plain
// @Success	302	{string}
// @Router		/auth/google [get]
func RedirectToGoogle(c *fiber.Ctx) error {
	state, err := oauth.CreateState()
	if err != nil {
		panic(err)
	}

	url := oauth.GetAuthCodeURL(state)
	cookie := new(fiber.Cookie)
	cookie.Name = "state"
	cookie.Value = state
	cookie.Expires = time.Now().Add(2 * time.Hour)

	c.Cookie(cookie)

	return c.Redirect(url)
}

// @Summary Handle callback from google signin
// @ID			callback-google
// @Produce	plain
// @Success	200	{string}
// @Router		/auth/google/callback [get]
func Callback(c *fiber.Ctx) error {
	stateCookie := c.Cookies("state", "")
	stateForm := c.FormValue("state")
	if stateCookie != stateForm {
		return c.SendStatus(502)
	}

	code := c.FormValue("code")
	result, err := oauth.HandleCallback(code)
	if err != nil {
		return c.SendStatus(502)
	}

	// response, err := json.Marshal(&result)
	// if err != nil {
	// 	return c.SendStatus(502)
	// }

	err = utils.SetKV(c, "Email", result.Email, 2)
	if err != nil {
		return c.SendStatus(502)
	}

	if controller.CheckUserInDB(result.Email) {
		utils.SetKV(c, "IsInDB", "true", 2)
	}
	return c.Redirect("/")
}

func AuthRoute(router fiber.Router) {
	router.Get("/google", RedirectToGoogle)

	router.Get("/google/callback", Callback)
}
