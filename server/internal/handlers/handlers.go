package handlers

import (
	"app/internal/database"
	"app/internal/repositories"
	"net/http"
)

type Handlers interface {
	GetAlbums(writer http.ResponseWriter, request *http.Request)
	GetAlbum(writer http.ResponseWriter, request *http.Request)
}

type handlers struct {
	albumService repositories.Album
}

func NewHandlers(database database.Gormw) Handlers {
	albumService := repositories.NewAlbum(database)
	return &handlers{
		albumService: albumService,
	}
}
