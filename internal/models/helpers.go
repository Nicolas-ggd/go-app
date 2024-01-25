package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", nil
	}

	hashedPassword := string(hash)

	return hashedPassword, nil
}

func CompareHashAndPasswordBcrypt(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return fmt.Errorf("incorrect password: %w", err)
	}

	return nil
}
