package models

type Listen struct {
	*Model
	AlbumID uint `gorm:"type:int" json:"album_id"`
	Album   Album `json:"-"`
}
