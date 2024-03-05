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
	CheckUserIsExist(id uint64) (bool, error)
}

type User struct {
	ID          uint64         `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Password    string         `json:"-"`
	PhoneNumber sql.NullString `json:"phone_number"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Token       []*Token       `json:"-" gorm:"foreignKey:UserID"`
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
	err := db.DB.Preload("Token").Scopes(EmailScope(email)).First(&us).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find user with email: %s", email)
	}

	return us, nil
}

func (us *User) CheckUserIsExist(id uint64) (bool, error) {
	var count int64

	if err := db.DB.Model(&User{}).Scopes(IdScope(id)).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
