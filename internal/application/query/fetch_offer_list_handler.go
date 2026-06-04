package query

import (
	application_shared "github.com/Ozenkol/rbk-go-final/internal/application/shared"
	"github.com/Ozenkol/rbk-go-final/internal/domain/offer"
)

type OfferFilter struct {
	ClientID string
}

type FetchOfferQuery struct {
	OfferID string
	Filter OfferFilter
	Pagination application_shared.Pagination
}

type FetchOfferHandler struct {
	offerRepository offer.OfferRepositoryInterface
}

func NewFetchOfferHandler(offerRepo offer.OfferRepositoryInterface) *FetchOfferHandler {
	return &FetchOfferHandler{offerRepository: offerRepo}
}

func (h *FetchOfferHandler) Handle(query FetchOfferQuery) ([]*offer.Offer, error) {
	if query.OfferID != "" {
		existingOffer, err := h.offerRepository.GetByID(query.OfferID)
		if err != nil {
			return nil, err
		}
		return []*offer.Offer{existingOffer}, nil
	}
	return []*offer.Offer{}, nil
}