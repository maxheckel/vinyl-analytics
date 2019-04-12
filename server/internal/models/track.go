package models

type Track struct {
	*Model
	Title    string `json:"title" gorm:"type:text"`
	Duration string `json:"duration" gorm:"type:text"`
	Position string   `json:"position" gorm:"type:text"`
	AlbumID  uint   `gorm:"type:int" json:"album_id"`
	Album    Album  `json:"-"`
}
