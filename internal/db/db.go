package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var DB *sql.DB

var (
	host   = os.Getenv("DB_HOST")
	port   = os.Getenv("DB_PORT")
	user   = os.Getenv("DB_USER")
	dbname = os.Getenv("DB_NAME")
	dbpass = os.Getenv("DB_PASSWORD")
	pgConn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", host, port, user, dbname)
)

func ConnectionDB() error {

	database, err := sql.Open("postgres", pgConn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Gracefully close the database connection on panic or program exit
	defer func() {
		if err := database.Close(); err != nil {
			fmt.Printf("Failed to close database connection: %v\n", err)
		}
	}()

	err = database.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	DB = database

	return nil
}

func DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s?sslmode=disable", user, dbpass, host, port)
}
