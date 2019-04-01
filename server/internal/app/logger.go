package app

import "github.com/sirupsen/logrus"

func BuildLogger() *logrus.Logger {
	return logrus.New()
}
