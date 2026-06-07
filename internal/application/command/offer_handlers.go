package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/offer"
)

type CreateOfferCommand struct {
	Offer *offer.Offer
}

type CreateOfferHandler struct {
	repo offer.OfferRepositoryInterface
}

func NewCreateOfferHandler(repo offer.OfferRepositoryInterface) *CreateOfferHandler {
	return &CreateOfferHandler{repo: repo}
}

func (h *CreateOfferHandler) Handle(ctx context.Context, cmd CreateOfferCommand) (*offer.Offer, error) {
	return h.repo.Create(cmd.Offer)
}

type UpdateOfferCommand struct {
	Offer *offer.Offer
}

type UpdateOfferHandler struct {
	repo offer.OfferRepositoryInterface
}

func NewUpdateOfferHandler(repo offer.OfferRepositoryInterface) *UpdateOfferHandler {
	return &UpdateOfferHandler{repo: repo}
}

func (h *UpdateOfferHandler) Handle(ctx context.Context, cmd UpdateOfferCommand) (*offer.Offer, error) {
	return h.repo.Update(cmd.Offer)
}

type DeleteOfferCommand struct {
	ID string
}

type DeleteOfferHandler struct {
	repo offer.OfferRepositoryInterface
}

func NewDeleteOfferHandler(repo offer.OfferRepositoryInterface) *DeleteOfferHandler {
	return &DeleteOfferHandler{repo: repo}
}

func (h *DeleteOfferHandler) Handle(ctx context.Context, cmd DeleteOfferCommand) error {
	return h.repo.Delete(cmd.ID)
}
