package offer

type OfferRepositoryInterface interface {
	Create(offer *Offer) error
	GetByID(id string) (*Offer, error)
	Update(offer *Offer) error
	Delete(id string) error
}