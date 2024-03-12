package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"log"
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

func ConnectionDB() *sql.DB {
	database, err := sql.Open("postgres", pgConn)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = database.Ping()
	if err != nil {
		log.Println(err)
		return nil
	}

	DB = database

	return database
}

func TestDatabaseConnection() error {
	// Open a database connection
	db, err := sql.Open("sqlite3", ":memory:?cache=shared")
	if err != nil {
		return err
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		return err
	}

	// Assign the database connection to the global variable
	DB = db

	return nil
}

func DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s?sslmode=disable", user, dbpass, host, port)
}
