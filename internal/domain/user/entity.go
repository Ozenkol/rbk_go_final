package user

import (
	"errors"
	"regexp"

	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"github.com/google/uuid"
)

type User struct {
	ID             string
	HumanName      shared.HumanName
	Email          string
	HashedPassword string
	IsVerified     bool
}

func NewUser(humanName shared.HumanName, email string, hashedPassword string) (User, error) {
	if humanName.FirstName == "" || humanName.LastName == "" {
		return User{}, errors.New("FirstName and LastName cannot be empty")
	}
	if email == "" {
		return User{}, errors.New("Email cannot be empty")
	}
	if email != "" && !isValidEmail(email) {
		return User{}, errors.New("Invalid email format")
	}
	if hashedPassword == "" {
		return User{}, errors.New("HashedPassword cannot be empty")
	}
	id := uuid.New().String() // "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	return User{
		ID:             id,
		HumanName:      humanName,
		Email:          email,
		HashedPassword: hashedPassword,
		IsVerified:     false,
	}, nil
}

func isValidEmail(email string) bool {
	// Simple regex for email validation
	const emailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailRegex, email)
	return matched
}

func (u *User) CheckPassword(password string) bool {
	return u.HashedPassword == hashPassword(password)
}

