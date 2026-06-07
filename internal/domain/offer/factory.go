package offer

type OfferFactoryInterface interface {
	CreateOffer(clientID string) (*Offer, error)
}

type OfferFactory struct {
	productRepository any
}

func NewOfferFactory(productRepository any) OfferFactoryInterface {
	return &OfferFactory{productRepository: productRepository}
}

func (f *OfferFactory) CreateOffer(clientID string) (*Offer, error) {
	return &Offer{
		ClientID: clientID,
	}, nil
}
