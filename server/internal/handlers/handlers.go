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
}

type handlers struct {
	albumService repositories.Album
	discogs discogs.Discogs
}

func NewHandlers(database database.Gormw, config *setup.Config) Handlers {
	discogsService := discogs.NewDiscogs(config.DiscogsToken)
	albumService := repositories.NewAlbum(database)

	return &handlers{
		albumService: albumService,
		discogs:discogsService,
	}
}
