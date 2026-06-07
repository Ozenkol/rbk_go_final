package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/user"
)

type DeleteUserCommand struct {
	ID string
}

type DeleteUserHandler struct {
	repo user.UserRepositoryInterface
}

func NewDeleteUserHandler(repo user.UserRepositoryInterface) *DeleteUserHandler {
	return &DeleteUserHandler{repo: repo}
}

func (h *DeleteUserHandler) Handle(ctx context.Context, cmd DeleteUserCommand) error {
	return h.repo.Delete(cmd.ID)
}
