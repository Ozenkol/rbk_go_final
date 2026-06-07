package note

type NoteRepositoryInterface interface {
	Create(note *Note) (*Note, error)
	GetByID(id string) (*Note, error)
	Update(note *Note) (*Note, error)
	Delete(id string) error
	List() ([]*Note, error)
}
