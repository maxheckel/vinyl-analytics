package app

import (
	"app/internal/global"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type App struct {
	router   *mux.Router
	Config   *Config
	database global.Gormw
	logger   *logrus.Logger
}

func NewApp() *App {
	app := &App{}
	app.buildLogger()

	cfg, err := BuildConfig()
	if err != nil {
		app.logger.Fatal(err)
	}
	app.Config = cfg
	database, err := BuildDatabase(cfg, app.logger)
	if err != nil {
		app.logger.Fatal(err)
	}
	app.database = database

	app.buildRoutes()


	return app
}


func (app *App) Serve() {
	app.validateSetup()
	app.logger.Info("Starting Server")
	app.logger.Fatal(http.ListenAndServe(":8081", app.router))
}


func (app *App) validateSetup() {
	if app.logger == nil {
		panic("nil logger when trying to serve")
	}
	if app.router == nil {
		app.logger.Fatal("nil router when trying to serve")
	}
	if app.database == nil {
		app.logger.Fatal("nil database when trying to serve")
	}
	if app.Config == nil {
		app.logger.Fatal("nil Config when trying to serve")
	}
}
