package models

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"os"
)

type Type string

const (
	Auth       Type = "auth"
	Validation Type = "validation"
)

//type Token struct {
//	ID     uint64 `gorm:"primaryKey"`
//	Hash   []byte
//	UserID uint64 `json:"id" gorm:"index:,unique,composite:idx_user_unique_token"`
//	Type   Type   `gorm:"index:,unique,composite:idx_user_unique_token"`
//}

type Token struct {
	ID     uint64 `json:"id"`
	Hash   []byte
	UserID uint64 `json:"user_id"`
	Type   Type   `json:"type"`
}

type TokenRepository struct {
	DBWrapper
}

// CreateJWT generates a JSON Web Token (JWT) containing user data.
//
// Parameters:
//
//	-UserId: The user ID to be included in the JWT claims.
//
// Returns:
//
//	-The signed JWT string, or an empty string if an error occurs.
//	-An error if there was a problem generating or signing the JWT.
func (tr *TokenRepository) CreateJWT(UserId uint64) (string, error) {
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

func (tr *TokenRepository) InsertToken(token *Token) error {
	//err := db.DB.Create(&Token{UserID: token.UserID, Hash: token.Hash, Type: token.Type}).Error
	//if err != nil {
	//	log.Printf("Error creating token: %v", err)
	//	return fmt.Errorf("failed to creating token: %v", err)
	//}
	query := `INSERT INTO user_tokens (hash, user_id, type)
	VALUES ($1, $2, $3)
	RETURNING id, hash, user_id, type`

	err := tr.DB.QueryRow(query, &token)
	if err != nil {
		log.Printf("Error creating token: %v", err)
		return fmt.Errorf("failed to creating token: %v", err)
	}

	return nil
}

func (tr *TokenRepository) DeleteToken(userId uint64) error {

	query := `UPDATE user_tokens SET deleted_at = $1
    WHERE id = $2`

	err := tr.DB.QueryRow(query, userId)
	if err != nil {
		log.Printf("Error deleting token: %v", err)
		return fmt.Errorf("failed to delete token: %v", err)
	}

	return nil
}

//func (tr *TokenRepository) CheckTokenExist(id uint64) (bool, error) {
//	var count int64
//
//	if err := db.DB.Model(&Token{}).Scopes(UserIdScope(id)).Count(&count).Error; err != nil {
//		return false, err
//	}
//
//	return count > 0, nil
//}
