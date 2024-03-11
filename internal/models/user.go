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
    WHERE users.email = $1`

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
	fmt.Println(user)
	return &user, nil
}

//func (us *UserRepository) CheckUserIsExist(id uint64) (bool, error) {
//	var count int64
//
//	if err := db.DB.Model(&User{}).Scopes(IdScope(id)).Count(&count).Error; err != nil {
//		return false, err
//	}
//
//	return count > 0, nil
//}
