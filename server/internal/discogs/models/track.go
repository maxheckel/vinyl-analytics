package models

type Track struct {
	Duration     string       `json:"duration"`
	Position     string       `json:"position"`
	Type         string       `json:"type_"`
	Title        string       `json:"title"`
	Extraartists []ArtistSlim `json:"extraartists,omitempty"`
}
