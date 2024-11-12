package main

import (
	"creepy/internal/service"
	"creepy/internal/storage/postgis"
	"creepy/pkg/config"
	"flag"
	"log"
	//"fmt"
	"os"
	"sync"
	"go.uber.org/zap"
)

var configPath = flag.String("config", ".env", "path to the configuration file")

func main() {
    // Initialize the database connection
    cfg := readConfig()
	//config.Set(cfg)
	log.Println(cfg)
	log.Fatal("out")
	db, err := postgis.NewPostgresGormConnection(cfg.DB)
    if err != nil {
        log.Fatal(err)
    }
    if err := postgis.AddExtension(db); err != nil {
        log.Fatal(err)
    }
	if err := postgis.Migrate(db); err != nil {
        log.Fatal(err)
    }


    // Initialize repositories
    propertyRepo := postgis.NewPropertyRepository(db)
    // Initialize other repositories...

    // Initialize services with repository interfaces
    propertyService := service.NewPropertyService(propertyRepo)
	_ = propertyService
    // Initialize other services...

    // Pass services to your handlers, bots, etc.
    // ...

	// Use the logger
	cfg.Logger.Info("Application starting", zap.String("version", "1.0.0"))

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		runCrawler(cfg)
	}()

	go func() {
		defer wg.Done()
		runTelegramBot(cfg)
	}()

	wg.Wait()

	cfg.Logger.Info("Application shutting down")
}

func readConfig() *config.Config {
	flag.Parse()

	if cfgPathEnv := os.Getenv("APP_CONFIG_PATH"); len(cfgPathEnv) > 0 {
		*configPath = cfgPathEnv
	}

	if len(*configPath) == 0 {
		log.Fatal("configuration file not found")
	}

	return config.NewConfig()
}

func runCrawler(cfg *config.Config) {
	cfg.Logger.Info("Starting property crawler")
	// Implement your crawling logic here
}

func runTelegramBot(cfg *config.Config) {
	cfg.Logger.Info("Starting Telegram bot")
	// Implement your Telegram bot logic here
}

/*
cfg.Logger.Info("User logged in",
    zap.String("username", "john_doe"),
    zap.Int("user_id", 12345))
*/
