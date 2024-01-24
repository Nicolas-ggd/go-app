package models

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
	"websocket/internal/db"
)

type User struct {
	ID          uint64         `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Password    string         `json:"-"`
	PhoneNumber sql.NullString `json:"phone_number"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func InsertUser(user *User) (*User, error) {
	hash, err := HashPassword(user.Password)
	if err != nil {
		log.Printf("Error generating password hash: %v", err)
		return nil, fmt.Errorf("failed to generate password hash: %v", err)
	}

	user.Password = hash

	err = db.DB.Create(&user).Error
	if err != nil {
		return nil, fmt.Errorf("fialed to insert user in databse")
	}

	return user, nil
}
