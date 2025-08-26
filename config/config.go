package config

type Config struct {
	App    App
	Server Server
	DB     DB
}

type App struct {
	Name string
}

type Server struct {
	Host string
	Port int
}

type DB struct {
	Username string
	Password string
	Host     string
	Port     int
	Name     string
}
