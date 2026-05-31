package client

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type Client struct {
	ID       string
	Person   shared.Person
	IsActive bool
}

func NewClient(person shared.Person, isActive bool) Client {
	return Client{
		Person:   person,
		IsActive: isActive,
	}
}
