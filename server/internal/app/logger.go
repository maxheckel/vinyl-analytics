package app

import "github.com/sirupsen/logrus"

func (app *App) buildLogger(){
	app.logger = logrus.New()
}