package discogs

import (
	"app/internal/discogs/services"
	"fmt"
	"gopkg.in/resty.v1"
)

type Discogs interface {
	Search() services.SearchService
}

type discogs struct {
	searchService services.SearchService
	client *resty.Client
}

func NewDiscogs(token string) *discogs {
	client := &resty.Client{}
	client.Header.Set("Authorization", fmt.Sprintf("Discogs token=%s", token))
	return &discogs{client: client}
}

func (d *discogs) Search() services.SearchService {
	if d.searchService == nil {
		searchService := services.NewSearchService(d.client)
		d.searchService = searchService
	}
	return d.searchService
}

