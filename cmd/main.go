package main

import (
	"creepy/internal/service"
	"creepy/internal/storage/mysql"
	"creepy/pkg/config"
	"flag"
	"log"
)

func main() {
    // Initialize the database connection
    cfg := readConfig()
	config.Set(cfg)

	db, err := mysql.NewMySQLGormConnection(cfg.DB)
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


var envFilePath = flag.String("envpath", "", ".env file path")

func readConfig() config.Config {
	flag.Parse()
	return config.ReadConfig(*envFilePath)
}
