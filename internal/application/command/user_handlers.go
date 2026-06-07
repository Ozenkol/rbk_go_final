package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/user"
)

type CreateUserCommand struct {
	User *user.User
}

type CreateUserHandler struct {
	repo    user.UserRepositoryInterface
	factory user.UserFactoryInterface
}

func NewCreateUserHandler(repo user.UserRepositoryInterface, factory user.UserFactoryInterface) *CreateUserHandler {
	return &CreateUserHandler{repo: repo, factory: factory}
}

func (h *CreateUserHandler) Handle(ctx context.Context, cmd CreateUserCommand) (*user.User, error) {
	u, err := h.factory.CreateUser(cmd.User.HumanName, cmd.User.Email, cmd.User.Password)
	if err != nil {
		return nil, err
	}
	return h.repo.Create(u)
}
