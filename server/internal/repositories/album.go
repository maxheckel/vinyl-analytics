package repositories

import (
	"app/internal/database"
	"app/internal/models"
)

type Album interface {
	FindAlbum(id int) (*models.Album, error)
	NewAlbum(name, artwork string, artist models.Artist) models.Album
	GetAlbums() ([]*models.Album, error)
}

type album struct {
	database database.Gormw
}

func NewAlbum(database database.Gormw) *album {
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

func (a *album) FindAlbum(id int) (*models.Album, error) {
	album := &models.Album{}
	err := a.database.First(&album, id).Error()
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (a *album) GetAlbums() ([]*models.Album, error){
	var albums []*models.Album
	err := a.database.Find(&albums).Error()
	if err != nil {
		return nil, err
	}
	return albums, nil
}