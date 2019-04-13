package repositories

import (
	"app/internal/database"
	"app/internal/models"
)

type Listen interface {
	CreateListen(albumId uint) (models.Listen, error)
	GetListens(albumId uint) ([]models.Listen, error)
	GetTotalListens(albumId uint) (int, error)
}

type listen struct {
	database database.Gormw
}

func NewListen(database database.Gormw) Listen {
	return &listen{database: database}
}

func (l *listen) CreateListen(albumId uint) (models.Listen, error) {
	listen := models.Listen{
		AlbumID:albumId,
	}
	err := l.database.Create(&listen).Error()
	return listen, err
}

func (l *listen) GetListens(albumId uint) ([]models.Listen, error){
	var listens []models.Listen
	err := l.database.Where("album_id = ?", albumId).Find(&listens).Error()
	return listens, err
}

func (l *listen) GetTotalListens(albumId uint) (int, error) {
	var count int
	err := l.database.Model(&models.Listen{}).Where("album_id = ?", albumId).Count(&count).Error()
	return count, err
}

