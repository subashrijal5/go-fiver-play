package route

import "github.com/gofiber/fiber/v2"

func SetWebRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{}, "layouts/default")
	})

	app.Get("/books", func(c *fiber.Ctx) error {
		return c.SendString("Get all books")
	})

	app.Get("/500", func(c *fiber.Ctx) error {
		return fiber.ErrInternalServerError
	})
}
