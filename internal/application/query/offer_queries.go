package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/offer"
)

type FetchOfferByID struct {
	ID string
}

type FetchOfferByIDHandler struct {
	repo offer.OfferRepositoryInterface
}

func NewFetchOfferByIDHandler(repo offer.OfferRepositoryInterface) *FetchOfferByIDHandler {
	return &FetchOfferByIDHandler{repo: repo}
}

func (h *FetchOfferByIDHandler) Handle(ctx context.Context, q FetchOfferByID) (*offer.Offer, error) {
	return h.repo.GetByID(q.ID)
}

type FetchOfferList struct{}

type FetchOfferListHandler struct {
	repo offer.OfferRepositoryInterface
}

func NewFetchOfferListHandler(repo offer.OfferRepositoryInterface) *FetchOfferListHandler {
	return &FetchOfferListHandler{repo: repo}
}

func (h *FetchOfferListHandler) Handle(ctx context.Context, q FetchOfferList) ([]*offer.Offer, error) {
	return h.repo.List()
}
