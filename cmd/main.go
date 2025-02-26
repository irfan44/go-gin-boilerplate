package main

import (
	"github.com/irfan44/go-gin-boilerplate/pkg/database"
	"log"

	"github.com/irfan44/go-gin-boilerplate/internal/config"
	"github.com/irfan44/go-gin-boilerplate/internal/server"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %s\n", err.Error())
	}
}

// @title Go REST API
// @version 1.0
// @description REST API using Golang
// @BasePath /
func main() {
	cfg := config.NewConfig()

	db, err := database.InitPGDB(cfg)
	if err != nil {
		return
	}

	s := server.NewServer(cfg, db)

	s.Run()
}
