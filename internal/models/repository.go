package models

import (
	"database/sql"
	"fmt"
)

// Repository provides DB of sql
type Repository struct {
	DB *sql.DB
}

// NewDBWrapper returns new wrapper with passed db parameters
func NewDBWrapper(db *sql.DB) *Repository {
	fmt.Println(db, "What is db?")
	return &Repository{
		DB: db,
	}
}
