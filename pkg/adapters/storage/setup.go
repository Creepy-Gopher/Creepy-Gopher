package storage

import (
	"creepy/config"
	"creepy/pkg/adapters/storage/entities"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMysqlGormConnection(dbConfig config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func Migrate(db *gorm.DB) {
	migrator := db.Migrator()

	migrator.AutoMigrate(&entities.FilterSet{})
}
