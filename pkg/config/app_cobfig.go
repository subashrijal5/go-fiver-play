package config

import (
	"os"
	"strconv"
	"time"
)

type App struct {
	Name  string
	Host  string
	Port  int
	Debug bool

	ReadTimeout time.Duration

	JwtSecret string
	JwtTtl    int
}

var app = &App{}

func AppConfig() *App {
	return app
}

func ConfigureAppConfig() {
	app.Name = os.Getenv("APP_NAME")
	app.Host = os.Getenv("APP_HOST")
	app.Port, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	app.Debug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))
	timeOut, _ := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
	app.ReadTimeout = time.Duration(timeOut) * time.Second

	app.JwtSecret = os.Getenv("APP_PORT")
	app.JwtTtl, _ = strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES"))
}
