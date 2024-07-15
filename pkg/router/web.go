package route

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/subashrijal5/go-fiber-boilerplate/app/controller/web"
)

func SetWebRoutes(app *fiber.App) {

	app.Get("/", controller.Index)

	app.Get("/features", controller.Features)

	app.Get("/pricing", controller.Pricing)

	app.Get("/contact", controller.Contact)

	app.Get("/500", func(c *fiber.Ctx) error {
		return fiber.ErrInternalServerError
	})

	auth := app.Group("/auth", func(c *fiber.Ctx) error {
		return c.Next()
	})

	auth.Get("/login", controller.LoginPage)
	// auth.Post("/login", controller.Login)
	auth.Get("/register", controller.RegisterPage)
	auth.Post("/register", controller.Register)

}
