package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/user"
)

type UpdateUserCommand struct {
	User *user.User
}

type UpdateUserHandler struct {
	repo user.UserRepositoryInterface
}

func NewUpdateUserHandler(repo user.UserRepositoryInterface) *UpdateUserHandler {
	return &UpdateUserHandler{repo: repo}
}

func (h *UpdateUserHandler) Handle(ctx context.Context, cmd UpdateUserCommand) (*user.User, error) {
	return h.repo.Update(cmd.User)
}
