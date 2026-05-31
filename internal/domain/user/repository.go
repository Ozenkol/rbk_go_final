package user

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type UserRepositoryInterface interface {
	Create(user *User) error
	GetByID(id string) (*User, error)
	Update(user *User) error
	Delete(id string) error

	LogIn(username, password string) (*User, error)
	LogOut(userID string) error
	Register(humanName shared.HumanName, email, password string) (*User, error)
	VerifyEmail(userID, token string) error
}