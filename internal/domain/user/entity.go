package user

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type User struct {
	ID        string
	HumanName shared.HumanName
	Email    string
	HashedPassword string
	IsVerified bool
}

func NewUser(humanName shared.HumanName, email string, hashedPassword string) User {
	if humanName.FirstName == "" || humanName.LastName == "" {
		panic("FirstName and LastName cannot be empty")
	}
	if email == "" {
		panic("Email cannot be empty")
	}
	if hashedPassword == "" {
		panic("HashedPassword cannot be empty")
	}
	return User{
		HumanName: humanName,
		Email:    email,
		HashedPassword: hashedPassword,
		IsVerified: false,
	}
}