package services

import (
	"app/internal/discogs/models"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/resty.v1"
)

type SearchService interface {
	Artists(query string, pagination *models.Pagination) ([]models.Artist, *models.Pagination, error)
}

type searchService struct {
	client *resty.Client
}

func NewSearchService(client *resty.Client) SearchService {
	return &searchService{client:client}
}

func (s *searchService) Artists(query string, pagination *models.Pagination) ([]models.Artist, *models.Pagination, error) {
	if pagination == nil{
		pagination = &models.Pagination{
			PerPage: 50,
			Page:    1,
		}
	}
	response, err := s.client.R().Get(fmt.Sprintf("https://api.discogs.com/database/search?type=artist&q=%s&page=%d&per_page=%d", query, pagination.Page, pagination.PerPage))
	if err != nil {
		return nil, nil, err
	}

	if response.Body() == nil {
		return nil, nil, errors.New("nil response body returned for request")
	}

	resultStruct := struct {
		Results []models.Artist `json:"results"`
		Pagination *models.Pagination `json:"pagination"`
	}{}
	err = json.Unmarshal(response.Body(), &resultStruct)

	if err != nil {
		return nil, nil, err
	}
	return resultStruct.Results, resultStruct.Pagination, nil
}

