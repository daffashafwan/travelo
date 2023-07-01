package utilities

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateJWT(userID, username string) string {
	// Define the JWT claims
	claims := jwt.MapClaims{
		"sub": userID,                          // Subject (User ID)
		"username": username,                   // Additional user data (optional)
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                // Issued at
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