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
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"-"`
	Token     []*Token  `json:"-"`
	IBank     []*IBank  `json:"bank_account"`
}

type UserForm struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type AuthUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
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

func (r *Repository) GetUserProfile(userId uint64) (*User, error) {
	var user User

	query := `SELECT id, created_at, name, email FROM users
	WHERE users.id = $1 AND users.deleted_at IS NULL`

	err := r.DB.QueryRow(query, userId).Scan(&user.ID, &user.CreatedAt, &user.Name, &user.Email)
	if err != nil {
		log.Printf("Error get user profile with ID: %v, got error: %v", userId, err)
		return nil, err
	}

	return &user, nil
}

func (r *Repository) UpdateProfile(userId uint64, user *User) (*User, error) {
	var u User

	query := `UPDATE users SET name = $2, updated_at = $3 WHERE users.id = $1 AND users.deleted_at IS NULL RETURNING id, name, email, updated_at`

	err := r.DB.QueryRow(query, userId, user.Name, time.Now()).Scan(&u.ID, &u.Name, &u.Email, &u.UpdatedAt)
	if err != nil {
		log.Printf("Can't update user profile with ID %v, got error: %v", userId, err)
		return nil, err
	}

	return &u, nil
}
