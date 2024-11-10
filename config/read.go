package config

import (
	"path/filepath"
)

// NewConfig initializes a new Config instance from environment variables
func NewConfig() *Config {
	return &Config{
		Server:   NewServerConfig(),
		DB:       NewDBConfig(),
		Telegram: NewTelegramConfig(),
	}
}

func readConfig(envPath string) *Config {

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
