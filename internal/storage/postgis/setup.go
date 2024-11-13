package postgis

import (
	"creepy/pkg/config"
	"creepy/internal/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresGormConnection(dbConfig config.DBConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func AddExtension(db *gorm.DB) error {
	return db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
}

func Migrate(db *gorm.DB) error {
	migrator := db.Migrator()
	err := migrator.AutoMigrate(
		&models.Filter{}, 
		&models.Property{}, 
		&models.Bookmark{},
		&models.UserSearchHistory{},
		&models.WatchList{},
		&models.User{},
	)
	if err != nil {
		return err
	}
	return nil
}