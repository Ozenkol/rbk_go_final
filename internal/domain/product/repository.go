package product

type ProductRepositoryInterface interface {
	GetProductByID(id string) (*Product, error)
}