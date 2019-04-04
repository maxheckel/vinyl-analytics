package app

import (
	"app/internal/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func (app *App) BuildRoutes() {
	r := mux.NewRouter()

	handlersProvider := handlers.NewHandlers(app.database, app.Config)
	r.HandleFunc("/albums", handlersProvider.GetAlbums).Methods(http.MethodGet)
	r.HandleFunc("/albums/{id}", handlersProvider.GetAlbum).Methods(http.MethodGet)
	r.HandleFunc("/albums", handlersProvider.CreateAlbum).Methods(http.MethodPost)
	r.HandleFunc("/albums/{id}", handlersProvider.DeleteAlbum).Methods(http.MethodDelete)

	r.HandleFunc("/discovery/albums/{name}", handlersProvider.DiscoverAlbums).Methods(http.MethodGet)

	r.HandleFunc("/listens/{id}", handlersProvider.CreateListen).Methods(http.MethodPost)
	r.HandleFunc("/listens/{id}", handlersProvider.GetListens).Methods(http.MethodGet)
	r.HandleFunc("/listens/{id}/count", handlersProvider.GetTotalListens).Methods(http.MethodGet)

	app.router = r
}
