package document

type DocumentRepositoryInterface interface {
	Create(document *Document) error
	GetByID(id string) (*Document, error)
	Update(document *Document) error
	Delete(id string) error
}
