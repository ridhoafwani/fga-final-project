package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your-secret-key")

// Claims represents JWT claims
type Claims struct {
	UserID uint `json:"userId"`
	jwt.StandardClaims
}

// GenerateJWT generates JWT token
func GenerateJWT(userID uint) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // Token valid for 7 days
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// VerifyJWT verifies JWT token
func VerifyJWT(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
}
