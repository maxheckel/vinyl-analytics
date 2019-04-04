package handlers

import (
	"app/internal/database"
	"app/internal/discogs"
	"app/internal/repositories"
	"app/internal/setup"
	"net/http"
)

type Handlers interface {
	GetAlbums(writer http.ResponseWriter, request *http.Request)
	GetAlbum(writer http.ResponseWriter, request *http.Request)
	DiscoverAlbums(writer http.ResponseWriter, request *http.Request)
	CreateAlbum(writer http.ResponseWriter, request *http.Request)
	DeleteAlbum(writer http.ResponseWriter, request *http.Request)

	CreateListen(writer http.ResponseWriter, request *http.Request)
	GetListens(writer http.ResponseWriter, request *http.Request)
	GetTotalListens(writer http.ResponseWriter, request *http.Request)
}

type handlers struct {
	albumService  repositories.Album
	listenService repositories.Listen
	discogs       discogs.Discogs
}

func NewHandlers(database database.Gormw, config *setup.Config) Handlers {
	discogsService := discogs.NewDiscogs(config.DiscogsToken)
	albumService := repositories.NewAlbum(database)
	listenService := repositories.NewListen(database)
	return &handlers{
		albumService:  albumService,
		discogs:       discogsService,
		listenService: listenService,
	}
}
