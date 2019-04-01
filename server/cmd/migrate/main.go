package main

import (
	"app/internal/setup"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := setup.BuildConfig()
	if err != nil {

	}
	logger := logrus.New()
	database, err := setup.BuildDatabase(cfg, logger)
	setup.Migrate(database, logger)
}
