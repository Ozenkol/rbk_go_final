package command

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/client"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
)

type CreateClientCommand struct {
	ID       string
	Person   shared.Person
	IsActive bool
}

type CreateClientHandler struct {
	clientRepo client.ClientRepositoryInterface
}

func NewCreateClientHandler(clientRepo client.ClientRepositoryInterface) *CreateClientHandler {
	return &CreateClientHandler{clientRepo: clientRepo}
}

func (h *CreateClientHandler) Handle(cmd CreateClientCommand) (string, error) {
	client := &client.Client{
		ID:       cmd.ID,
		Person:   cmd.Person,
		IsActive: cmd.IsActive,
	}
	savedClient, err := h.clientRepo.Save(client)
	if err != nil {
		return "", err
	}
	return savedClient.ID, nil
}