package models

import (
	"encoding/json"
	"net/url"
)

type AlbumSearch struct {
	Style       []string  `json:"style"`
	Barcode     []string  `json:"barcode"`
	Thumb       string    `json:"thumb"`
	Uri         string    `json:"uri"`
	Title       string    `json:"title"`
	Country     string    `json:"country"`
	Format      []string  `json:"format"`
	UserData    UserData  `json:"user_data"`
	Community   Community `json:"community"`
	Label       []string  `json:"label"`
	CoverImage  string    `json:"cover_image"`
	Catno       string    `json:"catno"`
	MasterURL   *url.URL  `json:"master_url"`
	Year        string    `json:"year"`
	Genre       []string  `json:"genre"`
	ResourceURL string    `json:"resource_url"`
	MasterId    uint      `json:"master_id"`
	Type        string    `json:"type"`
	ID          uint      `json:"id"`
}

func (a *AlbumSearch) UnmarshalJSON(data []byte) error {
	type Alias AlbumSearch
	aux := &struct {
		MasterURL string `json:"master_url"`
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
	return nil
}
