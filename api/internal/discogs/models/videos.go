package models


type Videos struct {
	Duration    int    `json:"duration"`
	Description string `json:"description"`
	Embed       bool   `json:"embed"`
	URI         string `json:"uri"`
	Title       string `json:"title"`
}
