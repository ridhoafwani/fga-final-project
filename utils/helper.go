package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetAuthenticatedUserId(c *gin.Context) (uint, error) {
	authToken := c.GetHeader("Authorization")
	tokenString := authToken[7:] // remove "Bearer " prefix from the token string

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKey
		}
		// Return the signing key
		return JWTKey, nil
	})
	if err != nil {
		return 0, err
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, jwt.ErrInvalidKey
	}

	// Retrieve the user ID from the claims
	userID, ok := claims["userId"].(float64)
	if !ok {
		return 0, jwt.ErrInvalidKey
	}

	// Convert user ID to uint
	return uint(userID), nil
}
