package main

import (
	"creepy/internal/bot"
	"creepy/internal/service"
	"creepy/pkg/config"
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go.uber.org/zap"
)

var configPath = flag.String("config", ".env", "path to the configuration file")

func main() {
	cfg := readConfig()

	app, err := service.NewAppContainer(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	cfg.Logger.Info("Application starting", zap.String("version", "1.0.0"))

	var wg sync.WaitGroup
	wg.Add(2)

	// Start the property crawler in a goroutine
	go func() {
		defer wg.Done()
		runCrawler(cfg, app)
	}()

	// Start the Telegram bot in a separate goroutine
	go func() {
		defer wg.Done()
		runTelegramBot(cfg, app, cfg.Logger)
	}()

	// Wait for exit signal
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit
	cfg.Logger.Info("Application shutting down")

	wg.Wait()
}

func readConfig() config.Config {
	flag.Parse()

	if cfgPathEnv := os.Getenv("APP_CONFIG_PATH"); len(cfgPathEnv) > 0 {
		*configPath = cfgPathEnv
	}

	if len(*configPath) == 0 {
		log.Fatal("configuration file not found")
	}

	return config.NewConfig()
}

func runCrawler(cfg config.Config, app *service.AppContainer) {
	cfg.Logger.Info("Starting property crawler")
	// Here, implement the real crawling logic
	// Keep this function active so that it doesn't terminate immediately
	for {
		// Example of running a periodic crawl task
		time.Sleep(30 * time.Minute)
		// Add crawling logic here
	}
}

func runTelegramBot(cfg config.Config, app *service.AppContainer, logger *zap.Logger) {
	logger.Info("Starting Telegram bot")

	// Initialize the bot with three arguments
	telegramBot, err := bot.NewBot(&cfg, app, logger)
	if err != nil {
		logger.Fatal("Failed to initialize Telegram bot", zap.Error(err))
	}

	// Start the bot
	telegramBot.Start()
}
