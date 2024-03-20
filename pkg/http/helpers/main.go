// Package helpers - helpers package is a http package helper, it contains different helper functions
// which is used to improve working in http package.
package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"websocket/internal/models"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// ExtractBearerToken func extract token from Authorization header
func ExtractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("invalid credentials, missing header, or bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 || jwtToken[0] != "Bearer" {
		return "", errors.New("incorrectly formatted or missing 'Bearer' keyword in authorization header")
	}

	return jwtToken[1], nil
}

// ValidateJWT function use private key to check that given jwt token is a valid signed jwt token or not
//
// Parameters:
//
//	-n(string): Parameter is jwt token as a string type
//
// Returns:
//
//	-Pointer of jwt.Token
//	-An error if given jwt token is not a signed with HMAC method, or it used a bad algorithm
func ValidateJWT(jwtToken string) (*jwt.Token, error) {
	dir, _ := os.Getwd()
	// keyPath := "/home/yamato/Documents/user-service"
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

// ParseJWTClaims func parse claims object
func ParseJWTClaims(header string) (*models.TokenClaim, error) {
	token, err := ExtractBearerToken(header)
	if err != nil {
		return nil, fmt.Errorf("failed to extract bearer token: %v", err)
	}

	parsedToken, err := ValidateJWT(token)
	if err != nil {
		return nil, fmt.Errorf("failed to validate JWT token: %v", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unexpected type for JWT claims")
	}

	var tokenClaim models.TokenClaim

	marshalled, err := json.Marshal(claims["user"])
	if err != nil {
		return nil, fmt.Errorf("marshalling error")
	}

	err = json.Unmarshal(marshalled, &tokenClaim)
	if err != nil {
		return nil, fmt.Errorf("error converting 'user' claim to float64")
	}

	return &tokenClaim, nil
}

// ParseObjectClaims func parse provided claims and get value according to passed parameter
func ParseObjectClaims(claims map[string]interface{}, key string) (int64, error) {
	valueFloat, ok := claims[key].(float64)
	if !ok {
		return 0, fmt.Errorf("value for key '%s' is not a valid float64", key)
	}

	value := int64(valueFloat)

	return value, nil
}

// GenerateRandomString generate random string with given "n" parameter, based on letterRunes(array of runs)
//
// Parameters:
//
//	-n(int): Parameter is int type, and it should be any range
//
// Returns:
//
//	-Generated random string
func GenerateRandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

// GetTokenClaim takes gin context and get user_claims from saved context
//
// Parameters:
//
//	-Context(Gin.Context): Pointer of gin.Context
//
// Returns:
//
//	-TokenClaim: model of TokenClaim
func GetTokenClaim(c *gin.Context) *models.TokenClaim {
	userObj, ok := c.Keys["user_claims"].(*models.TokenClaim)

	if !ok {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "user claims not found in context"})
		c.Abort()
		return nil
	}

	return userObj
}
