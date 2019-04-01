package repositories

import (
	"app/internal/models"
	"github.com/jinzhu/gorm"
)

type Album interface {
	FindAlbum(id int) *models.Album
	NewAlbum(name, artwork string, artist models.Artist) models.Album
}

type album struct {
	database *gorm.DB
}

func newAlbum(database *gorm.DB) *album {
	return &album{database: database}
}

func (a *album) NewAlbum(name, artwork string, artist models.Artist) models.Album {
	album := models.Album{
		Name:     name,
		Artwork:  artwork,
		ArtistID: artist.ID,
	}
	a.database.Create(&album)
	return album
}

func (a *album) FindAlbum(id int) *models.Album {
	album := &models.Album{}
	a.database.First(&album, id)
	return album
}


