package models

type Album struct {
	*Model
	Name      string   `gorm:"type:text" json:"name"`
	Artwork   string   `gorm:"type:text" json:"artwork"`
	Year      string   `gorm:"type:text" json:"year"`
	ArtistID  uint     `gorm:"type:int" json:"artist_id"`
	Artist    Artist   `json:"artist"`
	Tracks    []Track  `json:"tracks"`
	DiscogsId uint     `json:"discogs_id"`
	Listens   []Listen `json:"listens"`
}
