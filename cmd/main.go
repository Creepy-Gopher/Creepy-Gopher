package main

import (
    "log"
	"creepy/pkg/config"
    "creepy/internal/service"
    "creepy/internal/storage/mysql"
)

func main() {
    // Initialize the database connection
	config := config.Config{DB: config.DB{}}
	db, err := mysql.NewMySQLGormConnection(config.DB)
    if err != nil {
        log.Fatal(err)
    }
	mysql.Migrate(db)

    // Initialize repositories
    propertyRepo := mysql.NewMySQLPropertyRepository(db)
    // Initialize other repositories...

    // Initialize services with repository interfaces
    propertyService := service.NewPropertyService(propertyRepo)
	_ = propertyService
    // Initialize other services...

    // Pass services to your handlers, bots, etc.
    // ...
}
