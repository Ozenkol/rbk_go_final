package user

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type UserFactoryInterface interface {
	CreateUser(humanName shared.HumanName, email string, password string) (*User, error)
}

type UserFactory struct {
}

func NewUserFactory() UserFactoryInterface {
	return &UserFactory{}
}

func (f *UserFactory) CreateUser(humanName shared.HumanName, email string, password string) (*User, error) {
	
	hashedPassword := hashPassword(password)

	user := NewUser(humanName, email, hashedPassword)
	return &user, nil
	
}

func hashPassword(password string) string {
	// Implement a secure password hashing mechanism here, e.g., bcrypt
	// For demonstration purposes, we'll just return the password as is (not secure)
	return password
}