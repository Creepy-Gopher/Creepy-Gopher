package mysql

import (
	"creepy/pkg/config"
	"creepy/internal/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLGormConnection(dbConfig config.DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbConfig.Host, dbConfig.User, dbConfig.Pass, dbConfig.DBName, dbConfig.Port)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func Migrate(db *gorm.DB) {
	migrator := db.Migrator()
	migrator.AutoMigrate(
		&models.Filter{}, 
		&models.Property{}, 
		&models.User{},
		&models.Bookmark{},
		&models.UserSearchHistory{},
		&models.WatchList{},
	)
}