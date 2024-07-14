package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/subashrijal5/go-fiber-boilerplate/pkg/config"
	"github.com/subashrijal5/go-fiber-boilerplate/pkg/middleware"
	route "github.com/subashrijal5/go-fiber-boilerplate/pkg/router"
	"github.com/subashrijal5/go-fiber-boilerplate/utils/database"
	"github.com/subashrijal5/go-fiber-boilerplate/utils/logger"
)

func main() {

	config.Configure(".env")

	logger.SetUpLogger()
	lgr := logger.GetLogger()

	appConfig := config.AppConfig()

	// connect to database
	_, err := database.Connect()

	if err != nil {
		lgr.Panicf("failed database setup. error: %v", err)
		return
	}

	// Define Fiber config & app.
	fiberCfg := config.FiberConfig()
	app := fiber.New(fiberCfg)

	app.Static("/", "./public")

	// Middlewares
	middleware.DefaultMiddleware(app)

	// Routes
	route.SetApiRoutes(app)
	route.SetWebRoutes(app)

	// Start Server
	serverAddress := fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port)

	log.Fatal(app.Listen(serverAddress))
}
