package models

type Album struct {
	*Model
	Name     string `gorm:"type:text" json:"name"`
	Artwork  string `gorm:"type:text" json:"artwork"`
	ArtistID uint   `gorm:"type:int" json:"artist_id"`
	Artist   Artist `json:"artist"`
	DiscogsId uint `json:"discogs_id"`
}

