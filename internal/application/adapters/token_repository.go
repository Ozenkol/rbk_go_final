package adapters

type TokenRepositoryInterface interface {
	SetRefreshToken(userID string, token string) error
	ValidateRefreshToken(userID string, token string) (bool, error)
	DeleteRefreshToken(userID string) error
	DeleteUserRefreshTokens(userID string) error

	Save(userID string, token string) error
	Get(token string) (string, error)
}
