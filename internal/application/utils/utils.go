package utils

import (
	"fmt"
	"time"

	"github.com/Ozenkol/rbk-go-final/internal/domain/user"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your-secret-key")

func GenerateAccessToken(user *user.User) (string, error) {
	fmt.Printf("Generating token for user: %s\n", user.Email)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": user.ID,
			"sub": user.Email,
			"iss": "rbk",
			"exp": time.Now().Add(time.Hour * 72).Unix(),
			"iat": time.Now().Unix(),
		},
	)
	fmt.Printf("Claims for user %s: %v\n", user.Email, claims.Claims)
	token, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	fmt.Printf("Generated token for user %s: %s\n", user.Email, token)
	return token, nil
}