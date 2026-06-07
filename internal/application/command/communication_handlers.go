package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/communication"
)

type CreateCommunicationCommand struct {
	Communication *communication.Communication
}

type CreateCommunicationHandler struct {
	repo communication.CommunicationRepositoryInterface
}

func NewCreateCommunicationHandler(repo communication.CommunicationRepositoryInterface) *CreateCommunicationHandler {
	return &CreateCommunicationHandler{repo: repo}
}

func (h *CreateCommunicationHandler) Handle(ctx context.Context, cmd CreateCommunicationCommand) (*communication.Communication, error) {
	return h.repo.Create(cmd.Communication)
}

type UpdateCommunicationCommand struct {
	Communication *communication.Communication
}

type UpdateCommunicationHandler struct {
	repo communication.CommunicationRepositoryInterface
}

func NewUpdateCommunicationHandler(repo communication.CommunicationRepositoryInterface) *UpdateCommunicationHandler {
	return &UpdateCommunicationHandler{repo: repo}
}

func (h *UpdateCommunicationHandler) Handle(ctx context.Context, cmd UpdateCommunicationCommand) (*communication.Communication, error) {
	return h.repo.Update(cmd.Communication)
}

type DeleteCommunicationCommand struct {
	ID string
}

type DeleteCommunicationHandler struct {
	repo communication.CommunicationRepositoryInterface
}

func NewDeleteCommunicationHandler(repo communication.CommunicationRepositoryInterface) *DeleteCommunicationHandler {
	return &DeleteCommunicationHandler{repo: repo}
}

func (h *DeleteCommunicationHandler) Handle(ctx context.Context, cmd DeleteCommunicationCommand) error {
	return h.repo.Delete(cmd.ID)
}
