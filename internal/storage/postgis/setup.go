package postgis

import (
	"creepy/internal/models"
	"creepy/pkg/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresGormConnection initializes a new Gorm connection with provided DBConfig
func NewPostgresGormConnection(dbConfig config.DBConfig) (*gorm.DB, error) {
	// Format the DSN (Data Source Name) for PostgreSQL connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port)

	// Open a new Gorm connection to the PostgreSQL database
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// AddExtension adds necessary PostgreSQL extensions, like uuid-ossp, to the database
func AddExtension(db *gorm.DB) error {
	return db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
}

// Migrate runs automatic migrations for all models in the database
func Migrate(db *gorm.DB) error {
	migrator := db.Migrator()
	err := migrator.AutoMigrate(
		&models.User{},              // ابتدا User
		&models.Property{},          // سپس Property
		&models.Filter{},            // سپس Filter
		&models.Bookmark{},          // سپس Bookmark که به User و Property وابسته است
		&models.WatchList{},         // سپس WatchList که به User و Filter وابسته است
		&models.UserSearchHistory{}, // سپس UserSearchHistory که به User و Filter وابسته است
	)
	if err != nil {
		return err
	}
	return nil
}
