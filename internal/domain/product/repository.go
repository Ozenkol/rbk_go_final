package product

type ProductRepositoryInterface interface {
	Create(product *Product) (*Product, error)
	GetByID(id string) (*Product, error)
	Update(product *Product) (*Product, error)
	Delete(id string) error
	List() ([]*Product, error)
}
