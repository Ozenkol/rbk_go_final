package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/deal"
)

type CreateDealCommand struct {
	Deal *deal.Deal
}

type CreateDealHandler struct {
	repo deal.DealRepositoryInterface
}

func NewCreateDealHandler(repo deal.DealRepositoryInterface) *CreateDealHandler {
	return &CreateDealHandler{repo: repo}
}

func (h *CreateDealHandler) Handle(ctx context.Context, cmd CreateDealCommand) (*deal.Deal, error) {
	return h.repo.Create(cmd.Deal)
}

type UpdateDealCommand struct {
	Deal *deal.Deal
}

type UpdateDealHandler struct {
	repo deal.DealRepositoryInterface
}

func NewUpdateDealHandler(repo deal.DealRepositoryInterface) *UpdateDealHandler {
	return &UpdateDealHandler{repo: repo}
}

func (h *UpdateDealHandler) Handle(ctx context.Context, cmd UpdateDealCommand) (*deal.Deal, error) {
	return h.repo.Update(cmd.Deal)
}

type DeleteDealCommand struct {
	ID string
}

type DeleteDealHandler struct {
	repo deal.DealRepositoryInterface
}

func NewDeleteDealHandler(repo deal.DealRepositoryInterface) *DeleteDealHandler {
	return &DeleteDealHandler{repo: repo}
}

func (h *DeleteDealHandler) Handle(ctx context.Context, cmd DeleteDealCommand) error {
	return h.repo.Delete(cmd.ID)
}
