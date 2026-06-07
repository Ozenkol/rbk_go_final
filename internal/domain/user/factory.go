package user

import (
	"errors"

	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
)

type UserFactoryInterface interface {
	CreateUser(humanName shared.HumanName, email string, password string) (*User, error)
}

type UserFactory struct {
	userRepo UserRepositoryInterface
}

func NewUserFactory(userRepo UserRepositoryInterface) UserFactoryInterface {
	return &UserFactory{userRepo: userRepo}
}

func (f *UserFactory) CreateUser(humanName shared.HumanName, email string, password string) (*User, error) {

	existingUser, err := f.userRepo.FindByEmail(email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already in use")
	}	

	existingUser, err = f.userRepo.FindByHumanName(humanName)
	if err == nil && existingUser != nil {
		return nil, errors.New("human name already in use")
	}

	hashedPassword := hashPassword(password)

	user, err := NewUser(humanName, email, hashedPassword)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

