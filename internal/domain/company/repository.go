package company

type CompanyRepositoryInterface interface {
	Create(company *Company) (*Company, error)
	GetByID(id string) (*Company, error)
	Update(company *Company) (*Company, error)
	Delete(id string) error
	List() ([]*Company, error)
}
