package setting

type SettingRepositoryInterface interface {
	Create(setting *Setting) error
	GetByID(id string) (*Setting, error)
	Update(setting *Setting) error
	Delete(id string) error
}