package utilities

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(username string) string {
	// Define the JWT claims
	claims := JWTClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(), // Token expires in 1 hour
		},
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	secretKey := []byte("123456789")
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Failed to generate JWT:", err)
		return ""
	}

	return signedToken
}