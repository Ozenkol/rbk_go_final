package meeting

type MeetingRepositoryInterface interface {
	Save(meeting *Meeting) error
	FindByID(id string) (*Meeting, error)
	Update(meeting *Meeting) error
	Delete(id string) error
}