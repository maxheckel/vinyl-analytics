package handlers

import (
	"app/internal/discogs/models"
	"app/internal/responses"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *handlers) DiscoverAlbums(writer http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
	var pagination *models.Pagination
	if vars["page"] != "" || vars["per_page"] != "" {
		page, err := strconv.Atoi(vars["page"])
		perPage, err := strconv.Atoi(vars["per_page"])
		if err != nil {

		}
		pagination = &models.Pagination{
			PerPage: uint(perPage),
			Page:    uint(page),
		}
	}
	albums, pagination, err := h.discogs.Search().Albums(vars["name"], pagination)
	if err != nil {
		responses.NewIntegrationError(err.Error(), writer)
		return
	}

	searchResponses := struct {
		Albums []models.Album `json:"albums"`
		Pagination *models.Pagination `json:"pagination"`
	}{
		Albums: albums,
		Pagination: pagination,
	}
	responses.WriteResponse(200, searchResponses, writer)
}
