package deal

type DealRepositoryInterface interface {
	Create(deal *Deal) (*Deal, error)
	GetByID(id string) (*Deal, error)
	Update(deal *Deal) (*Deal, error)
	Delete(id string) error
	List() ([]*Deal, error)
}
