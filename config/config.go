package config

// Config holds the application configuration
type Config struct {
	Server   ServerConfig
	DB       DBConfig
	Telegram TelegramConfig
}

// ServerConfig holds the server configuration
type ServerConfig struct {
	Port                   string
	Host                   string
	TokenExpMinutes        string
	RefreshTokenExpMinutes string
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

// // NewConfig initializes a new Config instance from environment variables
// func NewConfig() *Config {
// 	return &Config{
// 		Server:   NewServerConfig(),
// 		DB:       NewDBConfig(),
// 		Telegram: NewTelegramConfig(),
// 	}
// }

