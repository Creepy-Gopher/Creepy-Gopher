package config

import (
	// "path/filepath"
	"os"
	"log"
	"github.com/joho/godotenv"
)

// NewConfig initializes a new Config instance from environment variables
// func NewConfig() *Config {
// 	return &Config{
// 		Server:   NewServerConfig(),
// 		DB:       NewDBConfig(),
// 		Telegram: NewTelegramConfig(),
// 	}
// }
// NewConfig initializes a new Config instance from environment variables
func NewConfig() *Config {
	return &Config{
		Server:   NewServerConfig(),
		DB:       NewDBConfig(),
		Telegram: NewTelegramConfig(),
	}
}

func ReadConfig(envPath string) *Config {
	err := godotenv.Load(envPath)	
 	if err != nil {
  		log.Fatalf("Error loading .env file: %s", err)
 	}

	return NewConfig()
}

// NewServerConfig initializes a new ServerConfig instance from environment variables
func NewServerConfig() ServerConfig {
	return ServerConfig{
		Port:					os.Getenv("SERVER_PORT"),
		Host:                   os.Getenv("SERVER_HOST"),
		TokenExpMinutes:        os.Getenv("TOKEN_EXP_MINUTES"),
		RefreshTokenExpMinutes: os.Getenv("REFRESH_TOKEN_EXP_MINUTES"),
		TokenSecret:            os.Getenv("TOKEN_SECRET"),
	}
}

// NewDBConfig initializes a new DBConfig instance from environment variables
func NewDBConfig() DBConfig {
	return DBConfig{
		User: 		os.Getenv("DB_USER"),
		Password: 	os.Getenv("DB_PASS"),
		Host: 		os.Getenv("DB_HOST"),
		Port: 		os.Getenv("DB_PORT"),
		Name: 		os.Getenv("DB_NAME"),
	}
}

// NewTelegramConfig initializes a new TelegramConfig instance from environment variables
func NewTelegramConfig() TelegramConfig {
	return TelegramConfig{
		BotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
	}
}
