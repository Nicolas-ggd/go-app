package models

import (
	"database/sql"
	"fmt"
)

// DBWrapper provides DB of sql
type DBWrapper struct {
	DB *sql.DB
}

// NewDBWrapper returns new wrapper with passed db parameters
func NewDBWrapper(db *sql.DB) *DBWrapper {
	fmt.Println(db, "What is db?")
	return &DBWrapper{
		DB: db,
	}
}
