package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/subashrijal5/go-fiber-boilerplate/pkg/config"
)

type Logger struct {
	*logrus.Logger
}

var logger = &Logger{}

// SetUpLogger settings
func SetUpLogger() {
	logger = &Logger{logrus.New()}
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetOutput(os.Stdout)

	if config.AppConfig().Debug {
		logger.SetLevel(logrus.DebugLevel)
	}
}

func GetLogger() *Logger {
	return logger
}
