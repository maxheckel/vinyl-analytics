package discogs

import (
	"app/internal/discogs/services"
	"fmt"
	"gopkg.in/resty.v1"
)

type Discogs interface {
	Search() services.SearchService
	Masters() services.MastersService
	Artist() services.ArtistService
}

type discogs struct {
	searchService services.SearchService
	mastersService services.MastersService
	artistService services.ArtistService
	client *resty.Client
}

func NewDiscogs(token string) *discogs {
	client := resty.New()
	client.SetHeader("Authorization", fmt.Sprintf("Discogs token=%s", token))
	return &discogs{client: client}
}

func (d *discogs) Search() services.SearchService {
	if d.searchService == nil {
		searchService := services.NewSearchService(d.client)
		d.searchService = searchService
	}
	return d.searchService
}


func (d *discogs) Masters() services.MastersService {
	if d.searchService == nil {
		mastersService := services.NewMasterService(d.client)
		d.mastersService = mastersService
	}
	return d.mastersService
}

func (d *discogs) Artist() services.ArtistService {
	if d.searchService == nil {
		artistService := services.NewArtistService(d.client)
		d.artistService = artistService
	}
	return d.artistService
}

