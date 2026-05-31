package offer

import "github.com/Ozenkol/rbk-go-final/internal/domain/product"

type OfferFactoryInterface interface {
	CreateOffer(userID, clientID, title, description string, productIDs []string) (*Offer, error)
}

type OfferFactory struct {
	productRepository product.ProductRepositoryInterface
}

func NewOfferFactory(productRepository product.ProductRepositoryInterface) OfferFactoryInterface {
	return &OfferFactory{
		productRepository: productRepository,
	}
}

func (f *OfferFactory) CreateOffer(userID, clientID, title, description string, productIDs []string) (*Offer, error) {
	
	offer := &Offer{
		ClientID:   clientID,
		DocumentID: "",
	}

	for _, productID := range productIDs {
		product, err := f.productRepository.GetProductByID(productID)
		if err != nil {
			return nil, err
		}
		offer.AddOfferItem(*product)
	}

	return offer, nil
}