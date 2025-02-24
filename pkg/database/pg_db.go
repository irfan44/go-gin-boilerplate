package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/irfan44/go-http-boilerplate/internal/config"
)

func InitPGDB(cfg config.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
	)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		fmt.Printf("Failed to ping to database: %s\n", err.Error())
		return nil, err
	}

	fmt.Println("Successfully connect to database")

	return db, nil
}
