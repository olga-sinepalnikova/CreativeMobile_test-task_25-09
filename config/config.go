package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
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
	logLevel := os.Getenv("LOG_LEVEL")
	switch strings.ToLower(logLevel) {
	case "debug":
		logrus.SetLevel(logrus.Level(5))
	default:
		logrus.SetLevel(logrus.Level(4))
	}

	logrus.Debug("Config.New")
	c := &Config{}
	err := godotenv.Load()
	if err != nil {
		logrus.Error("Error loading .env file")
		return nil, err
	}

	c.PostgresConn = os.Getenv("POSTGRES_CONN")
	c.PostgresPort = os.Getenv("POSTGRES_PORT")
	c.PostgresUser = os.Getenv("POSTGRES_USER")
	c.PostgresPass = os.Getenv("POSTGRES_PASSWORD")
	c.PostgresName = os.Getenv("POSTGRES_DATABASE")
	c.PostgresHost = os.Getenv("POSTGRES_HOST")

	return c, nil
}
