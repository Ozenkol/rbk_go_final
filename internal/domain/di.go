package domain

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/offer"
	"github.com/Ozenkol/rbk-go-final/internal/domain/product"
)

type diContainer struct {
	offerFactory *offer.OfferFactoryInterface
}

func newDIDomainContainer() *diContainer {
	return &diContainer{}
}

func (c *diContainer) CreateOfferFactory(productRepository product.ProductRepositoryInterface) offer.OfferFactoryInterface {
	return offer.NewOfferFactory(productRepository)
}