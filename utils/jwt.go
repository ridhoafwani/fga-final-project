// Package utils provides utility functions for handling JWT tokens.

package utils

import (
	"github.com/dgrijalva/jwt-go"
)

var JWTKey = "secret"

// GenerateJWT generates a JWT token for the given userID.
func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTKey))
}

// VerifyJWT verifies the validity of a JWT token and returns the userID.
func VerifyJWT(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, err
	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		return 0, err
	}
	return uint(userID), nil
}
