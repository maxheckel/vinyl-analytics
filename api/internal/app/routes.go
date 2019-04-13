package app

import (
	"app/internal/http_handlers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func (app *App) BuildRoutes() {
	r := mux.NewRouter()

	//// Setup common middleware
	cors := handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS", "PATCH", "DELETE"}),
		handlers.AllowedHeaders([]string{
			"Accept", "Accept-Language", "Content-Language", "Origin",
			"X-Requested-With", "Content-Type", "Authorization",
		}),
		handlers.AllowedOrigins([]string{"*"}),
	)

	r.Use(cors)
	handlersProvider := http_handlers.NewHandlers(app.database, app.Config)
	r.HandleFunc("/albums", handlersProvider.GetAlbums).Methods(http.MethodGet)
	r.HandleFunc("/albums/{id}", handlersProvider.GetAlbum).Methods(http.MethodGet)
	r.HandleFunc("/albums", handlersProvider.CreateAlbum).Methods(http.MethodPost, http.MethodOptions)
	//r.HandleFunc("/albums", func(writer http.ResponseWriter, request *http.Request) {
	//
	//}).Methods(http.MethodOptions)

	r.HandleFunc("/albums/{id}", handlersProvider.DeleteAlbum).Methods(http.MethodDelete)

	r.HandleFunc("/discovery/albums/{name}", handlersProvider.DiscoverAlbums).Methods(http.MethodGet)

	r.HandleFunc("/listens/{id}", handlersProvider.CreateListen).Methods(http.MethodPost)
	r.HandleFunc("/listens/{id}", handlersProvider.GetListens).Methods(http.MethodGet)
	r.HandleFunc("/listens/{id}/count", handlersProvider.GetTotalListens).Methods(http.MethodGet)

	r.HandleFunc("/listen", handlersProvider.Listen).Methods(http.MethodPost, http.MethodOptions)

	app.router = r
}
