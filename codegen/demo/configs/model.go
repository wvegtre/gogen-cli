package configs

type Conf struct {
	Server   Server
	Database Database
}

type Server struct {
	Port string
}

type Database struct {
	MySQL MySQL
}

type MySQL struct {
	User              string
	Host              string
	Port              int
	DB                string
	MaxLifetimeSecond int
	MaxIdleConns      int
}
