package file

type FileRepositoryInterface interface {
	Create(file *File) error
	GetByID(id string) (*File, error)
	Update(file *File) (*File, error)
	Delete(id string) error
}