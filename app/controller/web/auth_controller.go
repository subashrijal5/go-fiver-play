package web_controller

import "github.com/gofiber/fiber/v2"

func LoginPage(c *fiber.Ctx) error {
	return c.Render("pages/auth/login", fiber.Map{}, "layouts/default")
}

func RegisterPage(c *fiber.Ctx) error {
	return c.Render("pages/auth/register", fiber.Map{}, "layouts/default")
}

func Register(c *fiber.Ctx) error {
	// validate request
	// create user
	// send response
	return nil

}

func ForgotPasswordPage(c *fiber.Ctx) error {
	return c.Render("pages/auth/forgot-password", fiber.Map{}, "layouts/default")
}

func ResetPasswordPage(c *fiber.Ctx) error {
	return c.Render("pages/auth/reset-password", fiber.Map{}, "layouts/default")
}
