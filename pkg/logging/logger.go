package logging

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var once sync.Once
var logger *logrus.Logger

func GetLogger() *logrus.Logger {
	once.Do(func() {
		logger = logrus.New()

		logger.Out = os.Stdout
		logger.SetLevel(logrus.WarnLevel)

		logger.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
	})

	return logger
}
