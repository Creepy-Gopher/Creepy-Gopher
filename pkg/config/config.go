package config

import (
	"os"
	"strconv"

	"go.uber.org/zap"
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
	Port     string
	Name     string
}

// TelegramConfig holds the Telegram bot configuration
type TelegramConfig struct {
	BotToken string
}

// NewConfig initializes a new Config instance from environment variables
func NewConfig() Config {
	cfg := Config{
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
	//config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

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
	return DBConfig{
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASS", "123456"),
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		Name:     getEnv("DB_NAME", "magic_creeper"),
	}
}

// NewTelegramConfig initializes a new TelegramConfig instance from environment variables
func NewTelegramConfig() TelegramConfig {
	return TelegramConfig{
		BotToken: getEnv("TELEGRAM_BOT_TOKEN", "YOUR_TELEGRAM_BOT_TOKEN"),
	}
}

// Helper function to retrieve environment variables with a fallback default value
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Helper function to retrieve environment variables as an integer with a fallback default value
func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
