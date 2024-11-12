package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv" // افزودن پکیج godotenv برای بارگذاری .env
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config holds the application configuration
type Config struct {
	Server   ServerConfig
	DB       DBConfig
	Telegram TelegramConfig
	Logger   *zap.Logger
}

// ServerConfig holds the server configuration
type ServerConfig struct {
	Port                   string
	Host                   string
	TokenExpMinutes        int
	RefreshTokenExpMinutes int
	TokenSecret            string
}

// DBConfig holds the database configuration
type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
}

// TelegramConfig holds the Telegram bot configuration
type TelegramConfig struct {
	BotToken string
}

// NewConfig initializes a new Config instance from environment variables
func NewConfig() *Config {
	// Load environment variables from .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	cfg := &Config{
		Server:   NewServerConfig(),
		DB:       NewDBConfig(),
		Telegram: NewTelegramConfig(),
	}

	cfg.Logger = initLogger()
	return cfg
}

// initLogger initializes the Zap logger
func initLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

// NewServerConfig initializes a new ServerConfig instance from environment variables
func NewServerConfig() ServerConfig {
	return ServerConfig{
		Port:                   getEnv("SERVER_PORT", "8000"),
		Host:                   getEnv("SERVER_HOST", "0.0.0.0"),
		TokenExpMinutes:        getEnvAsInt("TOKEN_EXP_MINUTES", 1440),
		RefreshTokenExpMinutes: getEnvAsInt("REFRESH_TOKEN_EXP_MINUTES", 2880),
		TokenSecret:            getEnv("TOKEN_SECRET", "P@$$%Secret6677"),
	}
}

// NewDBConfig initializes a new DBConfig instance from environment variables
func NewDBConfig() DBConfig {
	portStr := getEnv("DB_PORT", "5432")
	port, err := strconv.Atoi(portStr) // Convert to int
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}
	return DBConfig{
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "123456"),
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     port,
		DBName:   getEnv("DB_NAME", "magic_creeper"),
	}
}

// NewTelegramConfig initializes a new TelegramConfig instance from environment variables
func NewTelegramConfig() TelegramConfig {
	return TelegramConfig{
		BotToken: getEnv("TELEGRAM_BOT_TOKEN", "YOUR_TELEGRAM_BOT_TOKEN"),
	}
}

// getEnv retrieves environment variables with a fallback default value
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt retrieves environment variables as an integer with a fallback default value
func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
