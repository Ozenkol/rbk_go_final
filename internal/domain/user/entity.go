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
	CompanyID	  string
	HashedPassword string
	Password	   string // Temporary for command to factory flow
	IsVerified     bool
	CreatedAt      int64
	UpdatedAt      int64
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
	id := uuid.New().String()
	return User{
		ID:             id,
		HumanName:      humanName,
		Email:          email,
		HashedPassword: hashedPassword,
		IsVerified:     false,
	}, nil
}

func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailRegex, email)
	return matched
}

func (u *User) CheckPassword(password string) bool {
	return u.HashedPassword == hashPassword(password)
}

func hashPassword(password string) string {
	return password
}
