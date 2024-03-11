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

type TokenClaim struct {
	UserId uint64 `json:"user_id"`
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
func (r *Repository) CreateJWT(UserID uint64) (string, error) {
	claims := &jwt.MapClaims{
		"ExpiresAt": 15000,
		"user": TokenClaim{
			UserId: UserID,
		},
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

func (r *Repository) InsertToken(token *Token) error {
	var tkn Token

	query := `INSERT INTO users_tokens (hash, user_id, type)
	VALUES ($1, $2, $3)
	RETURNING hash`

	err := r.DB.QueryRow(query, token.Hash, token.UserID, token.Type).Scan(&tkn.Hash)
	if err != nil {
		log.Printf("Error creating token: %v", err)
		return fmt.Errorf("failed to create token: %v", err)
	}

	return nil
}

func (r *Repository) DeleteToken(userId uint64) (bool, error) {
	query := `DELETE FROM users_tokens WHERE user_id = $1`

	res, err := r.DB.Exec(query, userId)
	if err == nil {
		_, err := res.RowsAffected()
		if err == nil {
			return true, nil
		}

	}

	return false, nil
}
