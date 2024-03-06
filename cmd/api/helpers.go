package api

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"os"
	"strings"
)

func (app *Application) extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("invalid credentials, missing header, or bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 || jwtToken[0] != "Bearer" {
		return "", errors.New("incorrectly formatted or missing 'Bearer' keyword in authorization header")
	}

	return jwtToken[1], nil
}

func (app *Application) ValidateJWT(jwtToken string) (*jwt.Token, error) {
	dir, _ := os.Getwd()
	secret, err := os.ReadFile(dir + "/tls/key.pem")
	if err != nil {
		log.Fatal(err)
	}

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}

func (app *Application) ParseJWTClaims(header string) (uint64, error) {
	token, err := app.extractBearerToken(header)
	if err != nil {
		return 0, fmt.Errorf("failed to extract bearer token: %v", err)
	}

	parsedToken, err := app.ValidateJWT(token)
	if err != nil {
		return 0, fmt.Errorf("failed to validate JWT token: %v", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("unexpected type for JWT claims")
	}

	userFloat, ok := claims["user"].(float64)
	if !ok {
		return 0, fmt.Errorf("error converting 'user' claim to float64")
	}

	userID := uint64(userFloat)

	return userID, nil
}
