package models

import (
	"app/internal/models"
	"encoding/json"
	"net/url"
)

type AlbumSearch struct {
	Style          []string               `json:"style"`
	Barcode        []string               `json:"barcode"`
	Thumb          string                 `json:"thumb"`
	URI            *url.URL               `json:"uri"`
	Title          string                 `json:"title"`
	Country        string                 `json:"country"`
	Format         []string               `json:"format"`
	UserData       UserData               `json:"user_data"`
	Community      Community              `json:"community"`
	Label          []string               `json:"label"`
	CoverImage     *url.URL               `json:"cover_image"`
	Catno          string                 `json:"catno"`
	MasterURL      *url.URL               `json:"master_url"`
	Year           string                 `json:"year"`
	Genre          []string               `json:"genre"`
	ResourceURL    string                 `json:"resource_url"`
	MasterId       uint                   `json:"master_id"`
	Type           string                 `json:"type"`
	ID             uint                   `json:"id"`
	ListenResponse *models.ListenResponse `json:"listen_response"`
	InCollection   bool                   `json:"in_collection"`
}

func (a *AlbumSearch) UnmarshalJSON(data []byte) error {
	type Alias AlbumSearch
	aux := &struct {
		MasterURL  string `json:"master_url"`
		CoverImage string `json:"cover_image"`
		URI        string `json:"uri"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	masterURL, err := url.Parse(aux.MasterURL)
	if err != nil {
		return err
	}
	a.MasterURL = masterURL

	coverImage, err := url.Parse(aux.CoverImage)
	if err != nil {
		return err
	}
	a.CoverImage = coverImage

	uri, err := url.Parse(aux.URI)
	if err != nil {
		return err
	}
	a.URI = uri
	return nil
}
