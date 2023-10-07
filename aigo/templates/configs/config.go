package configs

import "gen-templates/internal/app/config"

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
	MaxIdle           int
}

func LoadConf() Conf {
	c := Conf{}
	v := config.LoadConfig()
	if err := v.Unmarshal(&c); err != nil {
		panic(err)
	}
	return c
}
