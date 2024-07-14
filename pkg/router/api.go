package route

import "github.com/gofiber/fiber/v2"

func SetApiRoutes(app *fiber.App) {
	route := app.Group("/api")

	route.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

}
