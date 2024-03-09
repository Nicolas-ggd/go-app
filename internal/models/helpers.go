package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword securely hashes a given password using bcrypt.
//
// Parameters:
//
//	-password: The plain-text password to be hashed.
//
// Returns:
//
//	-The hashed password, or an empty string if an error occurs.
//	-An error if there was a problem generating the hash.
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", nil
	}

	hashedPassword := string(hash)

	return hashedPassword, nil
}

// CompareHashAndPasswordBcrypt compare hashed password and given password each other. function use bcrypt package to compare hash and password.
//
// Parameters:
//
//	-hash(string): String of hashed password
//	-password(string): String of password
//
// Returns:
//
//	-An error if hashed password and password doesn't match each other.
func CompareHashAndPasswordBcrypt(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return fmt.Errorf("incorrect password: %w", err)
	}

	return nil
}
