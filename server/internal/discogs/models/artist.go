package models

type ArtistSlim struct {
	Join        string `json:"join"`
	Name        string `json:"name"`
	Anv         string `json:"anv"`
	Tracks      string `json:"tracks"`
	Role        string `json:"role"`
	ResourceURL string `json:"resource_url"`
	ID          uint   `json:"id"`
}

type Artist struct {
	Profile        string   `json:"profile"`
	ReleasesURL    string   `json:"releases_url"`
	Name           string   `json:"name"`
	URI            string   `json:"uri"`
	Members        []Member `json:"members"`
	Urls           []string `json:"urls"`
	Images         []Image  `json:"images"`
	ResourceURL    string   `json:"resource_url"`
	ID             uint      `json:"id"`
	DataQuality    string   `json:"data_quality"`
	Namevariations []string `json:"namevariations"`
}
