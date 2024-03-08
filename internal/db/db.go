package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	host   = os.Getenv("DB_HOST")
	port   = os.Getenv("DB_PORT")
	user   = os.Getenv("DB_USER")
	dbname = os.Getenv("DB_NAME")
	dbpass = os.Getenv("DB_PASSWORD")
	pgConn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", host, port, user, dbname)
)

func ConnectionDB() (*sql.DB, error) {

	db, err := sql.Open("postgres", pgConn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Gracefully close the database connection on panic or program exit
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Printf("Failed to close database connection: %v\n", err)
		}
	}()

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s?sslmode=disable", user, dbpass, host, port)
}
