package web_controller

import (
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	return c.Render("pages/index", fiber.Map{}, "layouts/default")
}

func Features(c *fiber.Ctx) error {
	return c.Render("pages/features", fiber.Map{}, "layouts/default")
}

func Pricing(c *fiber.Ctx) error {
	return c.Render("pages/pricing", fiber.Map{}, "layouts/default")
}

func Contact(c *fiber.Ctx) error {
	return c.Render("pages/contact", fiber.Map{}, "layouts/default")
}
