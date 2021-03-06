package setup

import (
	"app/internal/database"
	"app/internal/models"
	"fmt"
	"github.com/sirupsen/logrus"
)

func BuildDatabase(config *Config, logger *logrus.Logger) (database.Gormw, error) {
	connection := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBName,
		config.DBPassword,
		config.DBSSLMode,
	)
	db, err := database.Openw("postgres", connection)
	if err != nil {
		return nil, err
	}
	logger.Info("Successfully connected to database")
	return db, nil
}

func Migrate(database database.Gormw, logger *logrus.Logger) {
	logger.Info("Starting migrations")
	database.AutoMigrate(models.Album{}, models.Artist{}, models.Listen{}, models.Track{})
	logger.Info("Migrations finished!")
}
