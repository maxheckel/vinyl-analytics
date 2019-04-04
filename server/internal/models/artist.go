package models

type Artist struct {
	*Model
	Name      string `gorm:"type:text" json:"name"`
	Bio       string `gorm:"type:text" json:"bio"`
	Image     string `gorm:"type:text" json:"image"`
	DiscogsId uint   `gorm:"type:int" json:"discogs_id"`
}
