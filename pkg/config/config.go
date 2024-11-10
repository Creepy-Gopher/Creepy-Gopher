package config

type Config struct {
	// Server Server: TODO
	DB     DB
}

func NewConfig(server Server, db DB) Config {
	return Config{
		// Server: server,
		DB:     db,
	}
}

type Server struct {
	Port                   string
	Host                   string
	TokenExpMinutes        uint64
	RefreshTokenExpMinutes uint64
	TokenSecret            string
}

func newServer(port string, host string, tokenExpMinutes uint64, refreshTokenExpMinutes uint64, tokenSecret string) *Server {
	return &Server{
		Port:                   port,
		Host:                   host,
		TokenExpMinutes:        tokenExpMinutes,
		RefreshTokenExpMinutes: refreshTokenExpMinutes,
		TokenSecret:            tokenSecret,
	}
}

type DB struct {
	User   string
	Pass   string
	Host   string
	Port   int
	DBName string
}

func newDB(user string, pass string, host string, port int, dbName string) *DB {
	return &DB{
		User:   user,
		Pass:   pass,
		Host:   host,
		Port:   port,
		DBName: dbName,
	}
}