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
	r.HandleFunc("/artists/search/{name}", handlersProvider.SearchArtist).Methods(http.MethodGet)
	app.router = r
}
