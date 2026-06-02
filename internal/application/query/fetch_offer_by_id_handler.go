package query

import "github.com/Ozenkol/rbk-go-final/internal/domain/offer"

type FetchOfferByIdQuery struct {
	OfferID string
}

type FetchOfferByIdHandler struct {
	offerRepository offer.OfferRepositoryInterface
}

func NewFetchOfferByIdHandler(offerRepository offer.OfferRepositoryInterface) *FetchOfferByIdHandler {
	return &FetchOfferByIdHandler{
		offerRepository: offerRepository,
	}
}

func (h *FetchOfferByIdHandler) Handle(query FetchOfferByIdQuery) (*offer.Offer, error) {
	return h.offerRepository.GetByID(query.OfferID)
}