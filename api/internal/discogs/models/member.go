package models


type Member struct {
	Active      bool   `json:"active"`
	ResourceURL string `json:"resource_url"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
}
