package config

import (
	"github.com/irfan44/go-gin-boilerplate/pkg/constants"
	"os"
)

type (
	Config struct {
		Http     httpConfig
		Postgres postgresConfig
		Jwt      jwtConfig
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

	jwtConfig struct {
		SecretKey string
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

		Jwt: jwtConfig{
			SecretKey: os.Getenv(constants.JwtSecretKey),
		},
	}

	return cfg
}
