package setup

import "github.com/sirupsen/logrus"

func BuildLogger() *logrus.Logger {
	return logrus.New()
}
