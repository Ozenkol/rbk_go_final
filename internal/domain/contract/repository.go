package contract

type ContractRepositoryInterface interface {
	Create(contract *Contract) (*Contract, error)
	GetByID(id string) (*Contract, error)
	Update(contract *Contract) (*Contract, error)
	Delete(id string) error
	List() ([]*Contract, error)
}
