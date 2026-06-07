package notification

type NotificationRepositoryInterface interface {
	Create(notification *Notification) (*Notification, error)
	GetByID(id string) (*Notification, error)
	Update(notification *Notification) (*Notification, error)
	Delete(id string) error
	List() ([]*Notification, error)
}
