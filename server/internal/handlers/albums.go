package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *handlers) GetAlbums(writer http.ResponseWriter, request *http.Request){
	
}

func (h *handlers) GetAlbum(writer http.ResponseWriter, request *http.Request)  {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {

	}
	album := h.albumService.FindAlbum(id)
	albumJson, err := json.Marshal(album)
	if err != nil {

	}
	writer.WriteHeader(200)
	_, _ = writer.Write(albumJson)
}