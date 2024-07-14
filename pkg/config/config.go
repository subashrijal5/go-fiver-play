package config

import (
	"log"
	"time"

	"github.com/gofiber/template/html/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Configure(envFile string) {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	ConfigureAppConfig()
	ConfigureDb()

}

// FiberConfig func for configuration Fiber app.
func FiberConfig() fiber.Config {

	// Create a new engine
	engine := html.New("./views", ".html")

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout:  time.Second * time.Duration(AppConfig().ReadTimeout),
		AppName:      AppConfig().Name,
		Views:        engine,
		ErrorHandler: DefaultErrorHandler,
	}
}

var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Define the path to error pages and layout
	errorPageName := "pages/errors/error"
	layoutName := "layouts/default"

	// Handle specific error codes
	switch code {
	case fiber.StatusInternalServerError, fiber.StatusNotFound:

		var message string
		if code == fiber.StatusInternalServerError {
			message = "Internal Server Error"
		} else {
			message = "Page not found"
		}

		// Render the error page with layout
		err = c.Status(code).Render(errorPageName, fiber.Map{
			"ErrorCode": code,
			"Message":   message,
		}, layoutName)
		if err != nil {
			log.Printf("Failed to render error page: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error: Failed to render error page")
		}
		return nil
	default:
		// For other error codes, log the error and send a generic message
		log.Printf("Unhandled error: %v", err)
		return c.Status(code).SendString("An error occurred")
	}
}
