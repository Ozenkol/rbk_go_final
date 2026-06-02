package query

import "github.com/Ozenkol/rbk-go-final/internal/domain/offer"

type FetchOfferQuery struct {
	OfferID string
}

type FetchOfferHandler struct {
	offerRepository offer.OfferRepositoryInterface
}