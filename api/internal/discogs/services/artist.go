package services

import (
	"app/internal/discogs/models"
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
)

type ArtistService interface {
	GetArtist(id uint) (*models.Artist, error)
}

type artistService struct {
	client *resty.Client
}

func NewArtistService(client *resty.Client) ArtistService {
	return &artistService{client: client}
}

func (m *artistService) GetArtist(id uint) (*models.Artist, error) {
	results, err := m.client.R().Get(fmt.Sprintf("https://api.discogs.com/artists/%d", id))
	if err != nil{
		return nil, err
	}
	artist := &models.Artist{}
	err = json.Unmarshal(results.Body(), artist)
	if err != nil {
		return nil, err
	}
	return artist, nil
}

