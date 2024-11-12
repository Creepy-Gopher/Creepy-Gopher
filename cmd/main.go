package main

import (
	"creepy/internal/bot"
	"creepy/internal/service"
	"creepy/internal/storage/postgis"
	"creepy/pkg/config"
	"flag"
	"log"
	"sync"

	"go.uber.org/zap"
)

var configPath = flag.String("config", ".env", "path to the configuration file")

func main() {
	// Initialize the configuration
	cfg := readConfig()
	config.Set(*cfg)

	// Set up the database connection
	db, err := postgis.NewPostgresGormConnection(cfg.DB)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Add necessary extensions and run migrations
	if err := postgis.AddExtension(db); err != nil {
		log.Fatal("Failed to add database extensions:", err)
	}
	if err := postgis.Migrate(db); err != nil {
		log.Fatal("Failed to run database migrations:", err)
	}

	// Initialize the repository and service layers
	propertyRepo := postgis.NewPropertyRepository(db)
	propertyService := service.NewPropertyService(propertyRepo)

	// Use the logger
	cfg.Logger.Info("Application starting", zap.String("version", "1.0.0"))

	// Run goroutines for crawler and Telegram bot
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		runCrawler(cfg)
	}()

	go func() {
		defer wg.Done()
		runTelegramBot(cfg, propertyService)
	}()

	wg.Wait()
	cfg.Logger.Info("Application shutting down")
}

func readConfig() *config.Config {
	flag.Parse()
	return config.NewConfig()
}

func runCrawler(cfg *config.Config) {
	cfg.Logger.Info("Starting property crawler")
	// پیاده‌سازی لاجیک کرالر
}

func runTelegramBot(cfg *config.Config, propertyService *service.PropertyService) {
	cfg.Logger.Info("Starting Telegram bot")

	// مثال ساده برای استفاده از پارامترها:
	telegramBot, err := bot.NewBot(cfg, propertyService, cfg.Logger)
	if err != nil {
		cfg.Logger.Fatal("Failed to create bot", zap.Error(err))
	}

	telegramBot.Start()
}
