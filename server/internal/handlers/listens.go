package handlers

import (
	"app/internal/models"
	"app/internal/responses"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *handlers) CreateListen(writer http.ResponseWriter, request *http.Request){
	album := h.getAlbumWithId(writer, request)
	if album == nil {
		return
	}
	listen, err := h.listenService.CreateListen(album.ID)
	if err != nil {
		responses.NewBadRequestErrror(err.Error(), writer)
		return
	}
	listen.Album = *album
	responses.WriteResponse(200, listen, writer)
}

func (h *handlers) GetListens(writer http.ResponseWriter, request *http.Request){
	album := h.getAlbumWithId(writer, request)
	if album == nil {
		return
	}
	listens, err := h.listenService.GetListens(album.ID)
	if err != nil {
		responses.NewBadRequestErrror(err.Error(), writer)
		return
	}
	for i := range listens{
		listens[i].Album = *album
	}
	responses.WriteResponse(200, listens, writer)
}
func (h *handlers) GetTotalListens(writer http.ResponseWriter, request *http.Request){
	album := h.getAlbumWithId(writer, request)
	if album == nil {
		return
	}
	count, err := h.listenService.GetTotalListens(album.ID)
	if err != nil {
		responses.NewBadRequestErrror(err.Error(), writer)
		return
	}
	responses.WriteResponse(200, struct {
		Listens int `json:"listens"`
	}{
		Listens:count,
	}, writer)
}

func (h *handlers) getAlbumWithId(writer http.ResponseWriter, request *http.Request) *models.Album{
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.NewBadRequestErrror("Could not convert id to int", writer)
		return nil
	}
	album, err := h.albumService.FindAlbum(uint(id))
	if err != nil {
		responses.NewBadRequestErrror("could not find album with error "+err.Error(), writer)
		return nil
	}
	return album
}