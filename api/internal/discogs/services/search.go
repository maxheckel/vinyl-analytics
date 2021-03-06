package services

import (
	"app/internal/discogs/models"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/resty.v1"
	"net/url"
)

//SearchService An interface for searching things in discogs' api
type SearchService interface {
	Artists(query string, pagination *models.Pagination) ([]models.ArtistSearch, *models.Pagination, error)
	Albums(query string, pagination *models.Pagination) ([]models.AlbumSearch, *models.Pagination, error)
	AlbumWithArtist(album, artist string, pagination *models.Pagination) ([]models.AlbumSearch, *models.Pagination, error)
}

type searchService struct {
	client *resty.Client
}

//NewSearchService returns a new instance of the search service
func NewSearchService(client *resty.Client) SearchService {
	return &searchService{client:client}
}

//ArtistSlim searches artists
func (s *searchService) Artists(query string, pagination *models.Pagination) ([]models.ArtistSearch, *models.Pagination, error) {
	response, err := s.performSearchWithPagination(query, "artist", pagination)
	resultStruct := struct {
		Results []models.ArtistSearch `json:"results"`
		Pagination *models.Pagination `json:"pagination"`
	}{}
	err = json.Unmarshal(response.Body(), &resultStruct)

	if err != nil {
		return nil, nil, err
	}
	return resultStruct.Results, resultStruct.Pagination, nil
}

//Albums searches albums
func (s *searchService) Albums(query string, pagination *models.Pagination) ([]models.AlbumSearch, *models.Pagination, error){
	response, err := s.performSearchWithPagination(query, "album", pagination)
	resultStruct := struct {
		Results []models.AlbumSearch  `json:"results"`
		Pagination *models.Pagination `json:"pagination"`
	}{}
	err = json.Unmarshal(response.Body(), &resultStruct)

	if err != nil {
		return nil, nil, err
	}
	return resultStruct.Results, resultStruct.Pagination, nil
}


func (s *searchService) performSearchWithPagination(query string, searchType string, pagination *models.Pagination) (*resty.Response, error){
	if pagination == nil{
		pagination = &models.Pagination{
			PerPage: 50,
			Page:    1,
		}
	}
	searchUrl := fmt.Sprintf("https://api.discogs.com/database/search?type=%s&q=%s&page=%d&per_page=%d", searchType, url.QueryEscape(query), pagination.Page, pagination.PerPage)
	response, err := s.client.R().Get(searchUrl)
	if err != nil {
		return nil, err
	}

	if response.Body() == nil {
		return nil, errors.New("nil response body returned for request")
	}
	return response, nil
}


func (s *searchService) AlbumWithArtist(album, artist string, pagination *models.Pagination) ([]models.AlbumSearch, *models.Pagination, error){
	if pagination == nil{
		pagination = &models.Pagination{
			PerPage: 50,
			Page:    1,
		}
	}
	searchUrl := fmt.Sprintf("https://api.discogs.com/database/search?type=album&artist=%s&q=%s&page=%d&per_page=%d", url.QueryEscape(artist), url.QueryEscape(album), pagination.Page, pagination.PerPage)
	fmt.Println(searchUrl)
	response, err := s.client.R().Get(searchUrl)
	if err != nil {
		return nil, nil, err
	}

	if response.Body() == nil {
		return nil, nil, errors.New("nil response body returned for request")
	}
	resultStruct := struct {
		Results []models.AlbumSearch  `json:"results"`
		Pagination *models.Pagination `json:"pagination"`
	}{}
	err = json.Unmarshal(response.Body(), &resultStruct)

	if err != nil {
		return nil, nil, err
	}
	return resultStruct.Results, resultStruct.Pagination, nil
}