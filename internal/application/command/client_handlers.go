package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/client"
)

type CreateClientCommand struct {
	Client *client.Client
}

type CreateClientHandler struct {
	repo client.ClientRepositoryInterface
}

func NewCreateClientHandler(repo client.ClientRepositoryInterface) *CreateClientHandler {
	return &CreateClientHandler{repo: repo}
}

func (h *CreateClientHandler) Handle(ctx context.Context, cmd CreateClientCommand) (*client.Client, error) {
	return h.repo.Create(cmd.Client)
}
