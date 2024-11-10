package postgis

import (
	"creepy/pkg/config"
	"creepy/internal/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresGormConnection(dbConfig config.DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		dbConfig.Host, dbConfig.User, dbConfig.Pass, dbConfig.DBName, dbConfig.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func Migrate(db *gorm.DB) error {
	migrator := db.Migrator()
	err := migrator.AutoMigrate(
		&models.Filter{}, 
		&models.Property{}, 
		&models.User{},
		&models.Bookmark{},
		&models.UserSearchHistory{},
		&models.WatchList{},
	)
	if err != nil {
		return err
	}
	return nil
}