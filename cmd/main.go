package main

import (
	"creepy_gopher/config"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
)

var configPath = flag.String("config", ".env", "path to the configuration file")

func main() {
	cfg := readConfig()

	fmt.Printf("Loaded Configuration: %+v\n", cfg)

	// Initialize your application components here
	// For example, you might want to create a crawler and a Telegram bot instance
	// app, err := service.NewAppContainer(cfg)
	// if err != nil {
	//     log.Fatal(err)
	// }

	var wg sync.WaitGroup
	wg.Add(2) // Adjust the number of goroutines based on your needs

	// Start the property crawler
	go func() {
		defer wg.Done()
		runCrawler(cfg) // This function should implement your crawling logic
	}()

	// Start the Telegram bot
	go func() {
		defer wg.Done()
		runTelegramBot(cfg) // This function should implement your Telegram bot logic
	}()

	wg.Wait() // Wait for all goroutines to finish
}

// readConfig reads the configuration from the specified path
func readConfig() *config.Config {
	flag.Parse()

	if cfgPathEnv := os.Getenv("APP_CONFIG_PATH"); len(cfgPathEnv) > 0 {
		*configPath = cfgPathEnv
	}

	if len(*configPath) == 0 {
		log.Fatal("configuration file not found")
	}

	cfg := config.NewConfig() // Initialize the config directly from environment variables

	return cfg
}

// runCrawler starts the property crawler using the provided configuration
func runCrawler(cfg *config.Config) {
	fmt.Println("Starting property crawler...")
	// Implement your crawling logic here using cfg
}

// runTelegramBot starts the Telegram bot using the provided configuration
func runTelegramBot(cfg *config.Config) {
	fmt.Println("Starting Telegram bot...")
	// Implement your Telegram bot logic here using cfg
}
