package config

import (
	"github.com/irfan44/go-http-boilerplate/pkg/constants"
	"os"
)

type (
	Config struct {
		Http     httpConfig
		Postgres postgresConfig
	}

	httpConfig struct {
		Port string
		Host string
	}

	postgresConfig struct {
		Port     string
		Host     string
		User     string
		Password string
		DBName   string
	}
)

func NewConfig() Config {
	cfg := Config{
		Http: httpConfig{
			Port: os.Getenv(constants.HTTPPort),
			Host: os.Getenv(constants.HTTPHost),
		},
		Postgres: postgresConfig{
			Port:     os.Getenv(constants.DBPort),
			Host:     os.Getenv(constants.DBHost),
			User:     os.Getenv(constants.DBUser),
			Password: os.Getenv(constants.DBPassword),
			DBName:   os.Getenv(constants.DBName),
		},
	}

	return cfg
}
