package main

import (
	"app/internal/app"
	"github.com/sirupsen/logrus"
)

func main(){
	cfg, err := app.BuildConfig()
	if err != nil {

	}
	logger := logrus.New()
	database, err := app.BuildDatabase(cfg, logger)
	app.Migrate(database, logger)
}
