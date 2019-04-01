package handlers

import (
	"app/internal/responses"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *handlers) GetAlbums(writer http.ResponseWriter, request *http.Request) {

}

func (h *handlers) GetAlbum(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {

	}
	album, err := h.albumService.FindAlbum(id)
	if err != nil {
		response := responses.NewNotFoundError("Could not find album with ID "+vars["id"])
		responses.WriteResponse(response.StatusCode, response, writer)
		return
	}

	responses.WriteResponse(200, album, writer)
}
