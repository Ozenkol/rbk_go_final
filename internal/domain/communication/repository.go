package communication

type CommunicationRepositoryInterface interface {
	Create(communication *Communication) (*Communication, error)
	GetByID(id string) (*Communication, error)
	Update(communication *Communication) (*Communication, error)
	Delete(id string) error
	List() ([]*Communication, error)
}
