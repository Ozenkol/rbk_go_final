package offer

type OfferRepositoryInterface interface {
	Create(offer *Offer) (*Offer, error)
	GetByID(id string) (*Offer, error)
	Update(offer *Offer) (*Offer, error)
	Delete(id string) error
	List() ([]*Offer, error)
}
