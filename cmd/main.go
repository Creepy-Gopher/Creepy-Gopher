package main

import (
	"creepy/internal/service"
	"creepy/internal/storage/postgis"
	"creepy/pkg/config"
	"flag"
	"log"
)

func main() {
    // Initialize the database connection
    cfg := readConfig()
	config.Set(cfg)

	db, err := postgis.NewPostgresGormConnection(cfg.DB)
    if err != nil {
        log.Fatal(err)
    }
	if err := postgis.Migrate(db); err != nil {
        log.Fatal(err)
    }

    // Initialize repositories
    propertyRepo := postgis.NewMySQLPropertyRepository(db)
    // Initialize other repositories...

    // Initialize services with repository interfaces
    propertyService := service.NewPropertyService(propertyRepo)
	_ = propertyService
    // Initialize other services...

    // Pass services to your handlers, bots, etc.
    // ...
}


var envFilePath = flag.String("envpath", "", ".env file path")

func readConfig() config.Config {
	flag.Parse()
	return config.ReadConfig(*envFilePath)
}
