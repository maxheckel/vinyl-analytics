package repositories

import (
	"app/internal/database"
	"app/internal/models"
	"errors"
	"fmt"
)

type Album interface {
	FindAlbum(id uint) (*models.Album, error)
	NewAlbum(name, artwork string, artist models.Artist, tracks []models.Track) models.Album
	GetAlbums() ([]*models.Album, error)
	CreateAlbum(newAlbum *models.Album) error
	DeleteAlbum(albumId uint) error
	GetByDiscogsId(discogsId uint) (*models.Album, error)
}

type album struct {
	database database.Gormw
}

func NewAlbum(database database.Gormw) *album {
	return &album{database: database}
}

func (a *album) NewAlbum(name, artwork string, artist models.Artist, tracks []models.Track) models.Album {
	album := models.Album{
		Name:     name,
		Artwork:  artwork,
		ArtistID: artist.ID,
		Tracks:   tracks,
	}
	a.database.Create(&album)
	return album
}

func (a *album) FindAlbum(id uint) (*models.Album, error) {
	album := &models.Album{}
	err := a.database.Preload("Artist").Preload("Tracks").Preload("Listens").First(&album, id).Error()
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (a *album) GetAlbums() ([]*models.Album, error) {
	var albums []*models.Album
	err := a.database.Preload("Artist").Find(&albums).Error()
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func (a *album) CreateAlbum(newAlbum *models.Album) error {
	var artistExists models.Artist
	a.database.Where("discogs_id = ?", newAlbum.Artist.DiscogsId).Find(&artistExists)
	if artistExists.ID > 0 {
		newAlbum.Artist = artistExists
		newAlbum.ArtistID = artistExists.ID
	}

	return a.database.Create(newAlbum).Error()
}

func (a *album) DeleteAlbum(albumId uint) error {
	var albumToDelete models.Album
	a.database.Preload("Artist").Find(&albumToDelete, albumId)
	if albumToDelete.ID == 0 {
		return errors.New(fmt.Sprintf("album with id %d does not exist", albumId))
	}
	var otherAlbum models.Album
	err := a.database.Where("artist_id = ?", albumToDelete.ArtistID).Find(&otherAlbum).Error()
	if otherAlbum.ID == 0 {
		err = a.database.Model(&albumToDelete.Artist).Delete(&albumToDelete.Artist).Error()
	}
	err = a.database.Model(&albumToDelete).Delete(&albumToDelete).Error()

	if err != nil {
		return err
	}
	return nil
}

func (a *album) GetByDiscogsId(discogsId uint) (*models.Album, error) {
	var albumFound *models.Album
	err := a.database.Where("discogs_id = ?", discogsId).Find(&albumFound).Error()
	if err != nil {
		return nil, err
	}
	if albumFound.ID > 0 {
		return albumFound, nil
	}
	return nil, errors.New("could not find album")
}
