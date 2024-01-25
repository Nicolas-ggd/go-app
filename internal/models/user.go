package models

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
	"websocket/internal/db"
)

type UserModelInterface interface {
	InsertUser(user *AuthUser) (*User, error)
	GetByEmail(email string) (*User, error)
}

type User struct {
	ID          uint64         `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Password    string         `json:"-"`
	PhoneNumber sql.NullString `json:"phone_number"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Token       []*Token       `json:"user_token" gorm:"foreignKey:UserID"`
}

type AuthUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (us *User) InsertUser(user *AuthUser) (*User, error) {
	hash, err := HashPassword(user.Password)
	if err != nil {
		log.Printf("Error generating password hash: %v", err)
		return nil, fmt.Errorf("failed to generate password hash: %v", err)
	}

	u := User{
		Email:    user.Email,
		Password: hash,
	}

	err = db.DB.Create(&u).Error
	if err != nil {
		return nil, fmt.Errorf("fialed to insert user in databse")
	}

	return &u, nil
}

func (us *User) GetByEmail(email string) (*User, error) {
	err := db.DB.Scopes(EmailScope(email)).First(&us).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find user with email: %s", email)
	}

	return us, nil
}
