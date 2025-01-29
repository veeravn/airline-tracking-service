package utils

import (
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func SetupLogger() {
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Logger.SetLevel(logrus.InfoLevel)
}
