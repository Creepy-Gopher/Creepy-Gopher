package config

// import (
// 	"log"
// 	"os"
// 	"strconv"

// 	"github.com/joho/godotenv"
// )

// func loadEnvVars(filePath string) {
// 	err := godotenv.Load(filePath)
// 	if err != nil {
// 		log.Fatalf("Error loading .env file: %v", err)
// 	}
// }

// func readServer() *Server {
// 	port := os.Getenv("SERVER_PORT")
// 	host := os.Getenv("SERVER_HOST")
// 	tokenExpMinutes, err := strconv.ParseUint(os.Getenv("TOKEN_EXP_MINUTES"), 10, 64)
// 	if err != nil {
// 		log.Fatal("TOKEN_EXP_MINUTES type is not valid.")
// 	}
// 	refreshTokenExpMinutes, err := strconv.ParseUint(os.Getenv("REFRESH_TOKEN_EXP_MINUTES"), 10, 64)
// 	if err != nil {
// 		log.Fatal("REFRESH_TOKEN_EXP_MINUTES type is not valid.")
// 	}
// 	secret := os.Getenv("TOKEN_SECRET")
// 	server := NewServerConfig(port, host, tokenExpMinutes, refreshTokenExpMinutes, secret)
// 	return server
// }

// func readDB() *DB {
// 	dbUser := os.Getenv("DB_USER")
// 	dbPass := os.Getenv("DB_PASS")
// 	dbHost := os.Getenv("DB_HOST")
// 	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
// 	if err != nil {
// 		log.Fatal("DB_PORT type is not valid.")
// 	}
// 	dbName := os.Getenv("DB_NAME")
// 	db := newDB(dbUser, dbPass, dbHost, dbPort, dbName)
// 	return db
// }

// func ReadConfig(envFilePath string) Config {
// 	loadEnvVars(envFilePath)
// 	server := readServer()
// 	db := readDB()
// 	config := NewConfig(*server, *db)
// 	return config
// }

// var cfg = Config{}

// func Set(config Config) {
// 	cfg = config
// }

// func Get() Config {
// 	return cfg
// }
