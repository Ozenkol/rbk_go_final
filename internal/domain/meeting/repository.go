package meeting

type MeetingRepositoryInterface interface {
	Create(meeting *Meeting) (*Meeting, error)
	GetByID(id string) (*Meeting, error)
	Update(meeting *Meeting) (*Meeting, error)
	Delete(id string) error
	List() ([]*Meeting, error)
}
