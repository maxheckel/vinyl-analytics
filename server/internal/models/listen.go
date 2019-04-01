package models

import "github.com/jinzhu/gorm"

type Listen struct {
	*gorm.Model
	AlbumID uint `gorm:"type:int"`
	Album   Album
}
