//go:build ignore

// Usage: go run scripts/gen_token.go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	secret := os.Getenv("APP_JWT_SECRET")
	if secret == "" {
		secret = "local-secret-change-in-production"
	}

	userID := "user-001"
	if len(os.Args) > 1 {
		userID = os.Args[1]
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	})

	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Println(signed)
}
