package db

import (
	"database/sql"
	"fmt"
	"junior/internal/config"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Config represents the database connection settings.
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// Connect establishes a connection to the PostgreSQL database.
func Connect(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
