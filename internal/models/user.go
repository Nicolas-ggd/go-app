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
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

type AuthUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *Repository) InsertUser(user *UserForm) (*User, error) {
	hash, err := HashPassword(user.Password)
	if err != nil {
		log.Printf("Error generating password hash: %v", err)
		return nil, fmt.Errorf("failed to generate password hash: %v", err)
	}

	stmt := `INSERT INTO users (name, email, password)
	VALUES ($1, $2, $3)`

	result, err := r.DB.Exec(stmt, user.Name, user.Email, hash)
	if err != nil {
		return nil, fmt.Errorf("fialed to insert user in databse: %v", err)
	}

	userID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last inserted ID: %v", err)
		return nil, fmt.Errorf("failed to get last inserted ID: %v", err)
	}

	u := &User{
		ID:    uint64(userID), // Convert int64 to int if applicable
		Name:  user.Name,
		Email: user.Email,
	}

	return u, nil
}

func (r *Repository) GetByEmail(email string) (*User, error) {
	//err := db.DB.Preload("Token").Scopes(EmailScope(email)).First(&us).Error
	//if err != nil {
	//	return nil, fmt.Errorf("failed to find user with email: %s", email)
	//}
	//
	//return us, nil
	var user User
	query := `SELECT id, created_at, name, email FROM users
    INNER JOIN token ON users.id = token.user_id
    WHERE users.id = $1`

	err := r.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Token,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to find user with email: %s", email)
	}

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
