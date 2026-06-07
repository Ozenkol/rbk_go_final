package query

import (
	application_shared "github.com/Ozenkol/rbk-go-final/internal/application/shared"
	"github.com/Ozenkol/rbk-go-final/internal/domain/offer"
)

type OfferFilter struct {
	ClientID string
}

type FetchOfferQuery struct {
	Filter OfferFilter
	Pagination application_shared.Pagination
}

type FetchOfferHandler struct {
	offerRepository offer.OfferRepositoryInterface
	offerQueryRepository OfferQueryRepositoryInterface
}

type OfferQueryRepositoryInterface interface {
	GetByID(id string) (*offer.Offer, error)
	FindByFilter(filter OfferFilter, pagination application_shared.Pagination) ([]*offer.Offer, error)
}

func NewFetchOfferHandler(offerRepo offer.OfferRepositoryInterface) *FetchOfferHandler {
	return &FetchOfferHandler{offerRepository: offerRepo}
}

func (h *FetchOfferHandler) Handle(query FetchOfferQuery) ([]*offer.Offer, error) {
	offers, err := h.offerQueryRepository.FindByFilter(query.Filter, query.Pagination)
	if err != nil {
		return nil, err
	}
	return offers, nil
}