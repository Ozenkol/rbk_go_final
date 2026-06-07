package service

import (
	"fmt"
	"time"

	"github.com/Ozenkol/rbk-go-final/internal/application/adapters"
	"github.com/Ozenkol/rbk-go-final/internal/domain/user"
	"github.com/golang-jwt/jwt/v5"
)



type AuthService struct {
	userRepository user.UserRepositoryInterface
	tokenRepository adapters.TokenRepositoryInterface
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

var secretKey = []byte("your-secret-key")

func NewAuthService(userRepository user.UserRepositoryInterface, tokenRepository adapters.TokenRepositoryInterface) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		tokenRepository: tokenRepository,
	}
}

func (s *AuthService) Authenticate(username, password string) (TokenPair, error) {
	user, err := s.userRepository.FindByEmail(username)
	if err != nil {
		return TokenPair{}, err
	}

	if !user.CheckPassword(password) {
		return TokenPair{}, fmt.Errorf("invalid credentials")
	}

	fmt.Printf("User %s authenticated successfully\n", user.Email)

	accessToken, err := generateAccessToken(user)

	refreshToken, err := generateRefreshToken(user) // Implement refresh token generation logic as needed

	jti, err := getRefreshIdFromToken(refreshToken)
	if err != nil {
		return TokenPair{}, err
	}

	err = s.tokenRepository.SetRefreshToken(user.ID, jti)
	if err != nil {
		return TokenPair{}, err
	}

	fmt.Printf("Generated token for user %s: %s\n", user.Email, accessToken)

	if err != nil {
		return TokenPair{}, err
	}

	return TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken, // Implement refresh token generation if needed
	}, nil
}

func (s *AuthService) GetUserByToken(token string) (string, error) {
	valid, err := validateToken(token)
	if err != nil {
		return "", err
	}
	if !valid {
		return "", err
	}
	userId, err := getUserIdFromToken(token)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (s *AuthService) ValidateAccessToken(accessToken string) (bool, error) {
	valid, err := validateToken(accessToken)
	if err != nil {
		return false, err
	}
	if !valid {
		return false, fmt.Errorf("invalid access token")
	}
	return true, nil
}

func generateRefreshToken(user *user.User) (string, error) {
	fmt.Printf("Generating refresh token for user: %s\n", user.Email)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": user.ID,
			"sub": user.Email,
			"jti": fmt.Sprintf("%d", time.Now().UnixNano()), // Unique identifier for the token
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
	fmt.Printf("Generated refresh token for user %s: %s\n", user.Email, token)
	return token, nil
}

func generateAccessToken(user *user.User) (string, error) {
	fmt.Printf("Generating refresh token for user: %s\n", user.Email)
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
	fmt.Printf("Generated refresh token for user %s: %s\n", user.Email, token)
	return token, nil
}

// Explanation
func validateToken(tokenString string) (bool, error) {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return secretKey, nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func getUserIdFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}
	jti, ok := claims["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("invalid refresh token ID")
	}
	return jti, nil
}

func getRefreshIdFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}
	jti, ok := claims["jti"].(string)
	if !ok {
		return "", fmt.Errorf("invalid refresh token ID")
	}
	return jti, nil
}