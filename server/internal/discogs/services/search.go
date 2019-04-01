package services

import (
	"app/internal/discogs/models"
	"gopkg.in/resty.v1"
)

type SearchService interface {
	SearchArtists(query string, pagination *models.Pagination) ([]models.Artist, error)
}

type searchService struct {
	client *resty.Client
}

func NewSearchService(client *resty.Client) SearchService {
	return &searchService{client:client}
}

func (s *searchService) SearchArtists(query string, pagination *models.Pagination) ([]models.Artist, error) {
	panic("implement me")
}

