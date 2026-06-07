package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/client"
)

type DeleteClientCommand struct {
	ID string
}

type DeleteClientHandler struct {
	repo client.ClientRepositoryInterface
}

func NewDeleteClientHandler(repo client.ClientRepositoryInterface) *DeleteClientHandler {
	return &DeleteClientHandler{repo: repo}
}

func (h *DeleteClientHandler) Handle(ctx context.Context, cmd DeleteClientCommand) error {
	return h.repo.Delete(cmd.ID)
}
