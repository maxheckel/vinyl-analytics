package models

import "github.com/jinzhu/gorm"

type Album struct {
	*gorm.Model
	Name     string `gorm:"type:text"`
	Artwork  string `gorm:"type:text"`
	ArtistID uint    `gorm:"type:int"`
	Artist   Artist
}
