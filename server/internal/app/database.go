package app

import (
	"app/internal/global"
	"app/internal/models"
	"fmt"
	"github.com/sirupsen/logrus"
)

func  BuildDatabase(config *Config, logger *logrus.Logger) (global.Gormw, error) {
	connection := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBName,
		config.DBPassword,
		config.DBSSLMode,
	)
	db, err := global.Openw("postgres", connection)
	if err != nil {
		return nil, err
	}
	logger.Info("Successfully connected to database")
	return db, nil
}

func Migrate(database global.Gormw, logger *logrus.Logger) {
	logger.Info("Starting migrations")
	database.AutoMigrate(&models.Album{}, &models.Artist{}, &models.Listen{})
	logger.Info("Migrations finished!")
}

