package command

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"github.com/subashrijal5/go-fiber-boilerplate/pkg/config"
	"github.com/subashrijal5/go-fiber-boilerplate/pkg/middleware"
	route "github.com/subashrijal5/go-fiber-boilerplate/pkg/router"
	"github.com/subashrijal5/go-fiber-boilerplate/utils/database"
	"github.com/subashrijal5/go-fiber-boilerplate/utils/logger"
)

var rootCommand = &cobra.Command{
	Use:   "Saas Cli",
	Short: "Saas Cli",
	Run: func(cmd *cobra.Command, args []string) {
		appConfig := config.AppConfig()
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
	},
}

func Execute() error {
	return rootCommand.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {

	config.Configure(".env")

	logger.SetUpLogger()
	lgr := logger.GetLogger()

	// connect to database
	err := database.Connect()

	if err != nil {
		lgr.Panicf("failed database setup. error: %v", err)
		return
	}
	// defer database.CloseDB()
}
