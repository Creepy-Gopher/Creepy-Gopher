package main

import (
	"creepy_gopher/config"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"go.uber.org/zap"
)

var configPath = flag.String("config", ".env", "path to the configuration file")

func main() {
	cfg := readConfig()

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
