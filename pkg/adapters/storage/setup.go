package storage

import (
	"creepy/config"
	"creepy/pkg/adapters/storage/entities"
	"fmt"

	_ "gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewPostgresGormConnection(dbConfig config.DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbConfig.Host, dbConfig.User, dbConfig.Pass, dbConfig.DBName, dbConfig.Port)
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
}

func Migrate(db *gorm.DB) {
	migrator := db.Migrator()

	migrator.AutoMigrate(&entities.Filter{})
}
