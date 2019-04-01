package handlers

import (
	"app/internal/global"
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

func NewHandlers(database global.Gormw) Handlers {
	return &handlers{

	}
}
