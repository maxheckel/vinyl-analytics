package http_handlers

import (
	"app/internal/models"
	"app/internal/responses"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *handlers) GetAlbums(writer http.ResponseWriter, request *http.Request) {
	albums, err := h.albumService.GetAlbums()
	if err != nil {
		responses.NewIntegrationError(err.Error(), writer)
		return
	}
	responses.WriteResponse(200, albums, writer)
}

func (h *handlers) GetAlbum(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.NewBadRequestErrror("Could not convert id to int", writer)
		return
	}
	album, err := h.albumService.FindAlbum(uint(id))
	if err != nil {
		responses.NewNotFoundError("Could not find album with ID "+vars["id"], writer)
		return
	}

	responses.WriteResponse(200, album, writer)
}

func (h *handlers) CreateAlbum(writer http.ResponseWriter, request *http.Request){
	var requestBody struct{
		DiscogsId uint `json:"discogs_id"`
	}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		responses.NewBadRequestErrror("could not parse request body", writer)
		return
	}
	if requestBody.DiscogsId <= 0 {
		responses.NewBadRequestErrror("invalid discogs id passed", writer)
		return
	}
	master, err := h.discogs.Masters().GetMaster(requestBody.DiscogsId)
	if err != nil {
		responses.NewIntegrationError(err.Error(), writer)
		return
	}
	artist, err := h.discogs.Artist().GetArtist(master.Artists[0].ID)
	if err != nil {
		responses.NewIntegrationError(err.Error(), writer)
		return
	}
	album := &models.Album{
		Name:     master.Title,
		Artwork:  master.Images[0].URI,
		Artist:   models.Artist{
			Name:  artist.Name,
			Bio:   artist.Profile,
			Image: artist.Images[0].URI,
			DiscogsId: artist.ID,
		},
	}
	err = h.albumService.CreateAlbum(album)
	if err != nil{
		responses.NewInternalError(err.Error(), writer)
		return
	}
	responses.WriteResponse(200, album, writer)
}

func (h *handlers) DeleteAlbum(writer http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.NewBadRequestErrror("Could not convert id to int", writer)
		return
	}
	err = h.albumService.DeleteAlbum(uint(id))
	if err != nil {
		responses.NewBadRequestErrror(err.Error(), writer)
		return
	}
	responses.WriteResponse(200, nil, writer)
}