package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/client"
)

type UpdateClientCommand struct {
	Client *client.Client
}

type UpdateClientHandler struct {
	repo client.ClientRepositoryInterface
}

func NewUpdateClientHandler(repo client.ClientRepositoryInterface) *UpdateClientHandler {
	return &UpdateClientHandler{repo: repo}
}

func (h *UpdateClientHandler) Handle(ctx context.Context, cmd UpdateClientCommand) (*client.Client, error) {
	return h.repo.Update(cmd.Client)
}
