package command

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"github.com/Ozenkol/rbk-go-final/internal/domain/user"
)

type CreateUserCommand struct {
	FirstName string
	LastName  string
	MiddleName string
	Email     string
	Password string
}

type CreateUserHandler struct {
	repo    user.UserRepositoryInterface
	factory user.UserFactoryInterface
}

func NewCreateUserHandler(repo user.UserRepositoryInterface, factory user.UserFactoryInterface) *CreateUserHandler {
	return &CreateUserHandler{
		repo:    repo,
		factory: factory,
	}
}

func (h *CreateUserHandler) Handle(cmd CreateUserCommand) (string, error) {
	humanName := shared.HumanName{
		FirstName: cmd.FirstName,
		LastName:  cmd.LastName,
		MiddleName: cmd.MiddleName,
	}
	password := cmd.Password
	user, err := h.factory.CreateUser(humanName, cmd.Email, password)
	if err != nil {
		return "", err
	}
	h.repo.Create(user)
	return user.ID, nil
}
