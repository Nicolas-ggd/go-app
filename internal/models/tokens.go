package models

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"os"
	"websocket/internal/db"
)

type Type string

const (
	Auth       Type = "auth"
	Validation Type = "validation"
)

type TokenModelInterface interface {
	CreateJWT(UserID uint64) (string, error)
	InsertToken(token *Token) error
	DeleteToken(userID uint64) error
	CheckTokenExist(id uint64) (bool, error)
}

type Token struct {
	ID     uint64 `gorm:"primaryKey"`
	Hash   []byte
	UserID uint64 `json:"id" gorm:"index:,unique,composite:idx_user_unique_token"`
	Type   Type   `gorm:"index:,unique,composite:idx_user_unique_token"`
}

func (t *Token) CreateJWT(UserId uint64) (string, error) {
	claims := &jwt.MapClaims{
		"ExpiresAt": 15000,
		"user":      UserId,
	}

	dir, _ := os.Getwd()
	content, err := os.ReadFile(dir + "/tls/key.pem")
	if err != nil {
		log.Fatal(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(content)
	if err != nil {
		fmt.Printf("Error signing JWT: %v", err)
		return "", err
	}

	return ss, nil
}

func (t *Token) InsertToken(token *Token) error {
	err := db.DB.Create(&Token{UserID: token.UserID, Hash: token.Hash, Type: token.Type}).Error
	if err != nil {
		log.Printf("Error creating token: %v", err)
		return fmt.Errorf("failed to creating token: %v", err)
	}

	return nil
}

func (t *Token) DeleteToken(userId uint64) error {
	var token Token

	result := db.DB.Unscoped().Scopes(UserIdScope(userId)).Delete(&token)
	if result.Error != nil {
		log.Printf("Error deleting token: %v", result.Error)

		return fmt.Errorf("failed to delete token: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("token with User_ID: %v not found", t.UserID)
	}

	return nil
}

func (t *Token) CheckTokenExist(id uint64) (bool, error) {
	var count int64

	if err := db.DB.Model(&Token{}).Scopes(UserIdScope(id)).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
