package models

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
	Token     []*Token  `json:"-"`
}

type UserForm struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type AuthUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *Repository) InsertUser(user *UserForm) (*User, error) {
	var usr User
	hash, err := HashPassword(user.Password)
	if err != nil {
		log.Printf("Error generating password hash: %v", err)
		return nil, fmt.Errorf("failed to generate password hash: %v", err)
	}

	stmt := `INSERT INTO users (name, email, password)
	VALUES ($1, $2, $3)
	RETURNING id, name, email, created_at`

	err = r.DB.QueryRow(stmt, user.Name, user.Email, hash).Scan(&usr.ID, &usr.Email, &usr.Name, &usr.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("fialed to insert user in databse: %v", err)
	}

	return &usr, nil
}

func (r *Repository) GetByEmail(email string) (*User, error) {
	var user User
	query := `SELECT id, created_at, name, email, password FROM users
    WHERE users.email = $1 AND users.deleted_at IS  NULL`

	err := r.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to find user with email: %s", email)
	}

	return &user, nil
}

func (r *Repository) DeleteAccount(userId uint64) error {
	query := `UPDATE users SET deleted_at = $1
	WHERE users.id = $2`

	_, err := r.DB.Exec(query, time.Now(), userId)
	if err != nil {
		log.Printf("failed to delete user account: %v", err)
		return err
	}

	return nil
}

func (r *Repository) RecoverAccount(userId uint64) error {
	query := `UPDATE users SET deleted_at = NULL
	WHERE users.id = $1`

	_, err := r.DB.Exec(query, userId)
	if err != nil {
		log.Printf("failed to delete user account: %v", err)
		return err
	}

	return nil
}
