package command

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"github.com/Ozenkol/rbk-go-final/internal/domain/user"
)

type CreateUserCommand struct {
	Name  string
	Email string
}

type CreateUserHandler struct {
	repo user.UserRepositoryInterface
	factory user.UserFactoryInterface
}

func NewCreateUserHandler(repo user.UserRepositoryInterface, factory user.UserFactoryInterface) *CreateUserHandler {
	return &CreateUserHandler{
		repo: repo,
		factory: factory,
	}
}	
	

func (h *CreateUserHandler) Handle(cmd CreateUserCommand) error {
	humanName := shared.HumanName{
		FirstName: cmd.Name,
		LastName:  "",
	}
	password := "default_password"
	user, err := h.factory.CreateUser(humanName, cmd.Email, password)
	if err != nil {
		return err
	}
	return h.repo.Create(user)
}