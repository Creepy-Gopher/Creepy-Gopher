package config

import (
	"log"

	"github.com/joho/godotenv"
)

// Load environment variables from a file
func loadEnvVars(filePath string) {
	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// Read server configuration
func readServer() ServerConfig {
	return NewServerConfig() // استفاده از تابع NewServerConfig برای مقداردهی ServerConfig
}

// Read database configuration
func readDB() DBConfig {
	return NewDBConfig() // استفاده از تابع NewDBConfig برای مقداردهی DBConfig
}

// ReadConfig initializes and returns the main config
func ReadConfig(envFilePath string) Config {
	loadEnvVars(envFilePath)
	server := readServer()
	db := readDB()
	config := NewConfig() // Initialize a new config
	config.Server = server
	config.DB = db
	return *config // بازگرداندن مقدار config بدون اشاره‌گر
}

// Global configuration variable
var cfg = Config{}

// Set sets the global configuration
func Set(config Config) {
	cfg = config
}

// Get retrieves the global configuration
func Get() Config {
	return cfg
}
