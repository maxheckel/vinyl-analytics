package models

type Album struct {
	Styles               []string     `json:"styles"`
	Genres               []string     `json:"genres"`
	Videos               []Videos     `json:"videos"`
	NumForSale           int          `json:"num_for_sale"`
	Title                string       `json:"title"`
	MostRecentRelease    int          `json:"most_recent_release"`
	MainRelease          int          `json:"main_release"`
	MainReleaseURL       string       `json:"main_release_url"`
	URI                  string       `json:"uri"`
	Artists              []ArtistSlim `json:"artists"`
	VersionsURL          string       `json:"versions_url"`
	DataQuality          string       `json:"data_quality"`
	MostRecentReleaseURL string       `json:"most_recent_release_url"`
	Year                 int          `json:"year"`
	Images               []Image      `json:"images"`
	ResourceURL          string       `json:"resource_url"`
	LowestPrice          float64      `json:"lowest_price"`
	ID                   int          `json:"id"`
	Tracklist            []Track      `json:"tracklist"`
}
