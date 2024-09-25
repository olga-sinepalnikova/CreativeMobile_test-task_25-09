package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	PostgresConn string
	PostgresPort string
	PostgresUser string
	PostgresPass string
	PostgresName string
	PostgresHost string
}

func New() (*Config, error) {
	c := &Config{}
	err := godotenv.Load()
	if err != nil {
		// todo: use logrus
		return nil, err
	}
	c.PostgresConn = os.Getenv("POSTGRES_CONNECTION")
	c.PostgresPort = os.Getenv("POSTGRES_PORT")
	c.PostgresUser = os.Getenv("POSTGRES_USER")
	c.PostgresPass = os.Getenv("POSTGRES_PASSWORD")
	c.PostgresName = os.Getenv("POSTGRES_DATABASE")
	c.PostgresHost = os.Getenv("POSTGRES_HOST")
	// todo: use logrus
	return c, nil
}
