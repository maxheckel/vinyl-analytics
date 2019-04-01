package handlers

import (
	"app/internal/discogs/models"
	"app/internal/responses"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *handlers) SearchArtist(writer http.ResponseWriter, request *http.Request){
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
	artists, pagination, err := h.discogs.Search().Artists(vars["name"], pagination)
	if err != nil {
		responses.NewIntegrationError(err.Error(), writer)
		return
	}

	searchResponses := struct {
		Artists    []models.Artist    `json:"artists"`
		Pagination *models.Pagination `json:"pagination"`
	}{
		Artists:    artists,
		Pagination: pagination,
	}
	fmt.Println(searchResponses.Artists)
	responses.WriteResponse(200, searchResponses, writer)
}