package tag

type TagRepositoryInterface interface {
	Create(tag *Tag) error
	GetByID(id string) (*Tag, error)
	Update(tag *Tag) (*Tag, error)
	Delete(id string) error
}