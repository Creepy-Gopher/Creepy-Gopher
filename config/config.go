package config

type Config struct {
	DB DB
}

type DB struct {
	Host   string
	Port   uint
	DBName string
	User   string
	Pass   string
}
