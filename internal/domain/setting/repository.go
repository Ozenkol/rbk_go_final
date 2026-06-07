package setting

type SettingRepositoryInterface interface {
	Create(setting *Setting) (*Setting, error)
	GetByID(id string) (*Setting, error)
	Update(setting *Setting) (*Setting, error)
	Delete(id string) error
	List() ([]*Setting, error)
}
