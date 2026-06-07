package persistence

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/application/adapters"
	"github.com/go-redis/redis/v8"
)

type Token struct {
	UserID string
	Token  string
}

type TokenRepository struct {
	rdb *redis.Client
}

func NewTokenRepository(rdb *redis.Client) (adapters.TokenRepositoryInterface, error) {
	return &TokenRepository{rdb: rdb}, nil
}

func (r *TokenRepository) SetRefreshToken(userID, token string) error {
	return r.rdb.Set(r.rdb.Context(), userID, token, 0).Err()
}

func (r *TokenRepository) ValidateRefreshToken(userID, token string) (bool, error) {
	storedToken, err := r.rdb.Get(r.rdb.Context(), userID).Result()
	if err != nil {
		return false, err
	}
	return storedToken == token, nil
}

func (r *TokenRepository) DeleteRefreshToken(userID string) error {
	return r.rdb.Del(r.rdb.Context(), userID).Err()
}

func (r *TokenRepository) DeleteUserRefreshTokens(userID string) error {
	return r.rdb.Del(r.rdb.Context(), userID).Err()
}

func (r *TokenRepository) Save(userID string, token string) error {
	return r.rdb.Set(context.Background(), token, userID, 0).Err()
}

func (r *TokenRepository) Get(token string) (string, error) {
	return r.rdb.Get(context.Background(), token).Result()
}
