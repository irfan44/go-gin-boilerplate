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
		log.Printf("Failed to ping to database: %s\n", err.Error())
		return nil, err
	}

	log.Println("Successfully connect to database")

	return db, nil
}

func InitializeTable(db *sql.DB) error {
	// TODO: fill init table query
	q1 := `
		CREATE TABLE IF NOT EXISTS users (
		    id SERIAL primary key,
		    username VARCHAR (255) UNIQUE NOT NULL,
		    password VARCHAR (255) NOT NULL,
		    role VARCHAR (30) NOT NULL CHECK (role IN ('TELLER', 'CUSTOMER', 'ADMIN')) DEFAULT 'CUSTOMER',
		    created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW()
		)
	`

	if _, err := db.Exec(q1); err != nil {
		log.Printf("Initialize table error: %s\n", err.Error())
		return err
	}

	log.Println("Successfully initiate table")

	return nil
}
