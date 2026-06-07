package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/setting"
	"gorm.io/gorm"
)

type SettingModel struct {
	ID   string `gorm:"primaryKey"`
}

type SettingRepository struct {
	db *gorm.DB
}

func NewSettingRepository(db *gorm.DB) (setting.SettingRepositoryInterface, error) {
	if err := db.AutoMigrate(&SettingModel{}); err != nil {
		return nil, err
	}
	return &SettingRepository{db: db}, nil
}

func (r *SettingRepository) Create(s *setting.Setting) (*setting.Setting, error) {
	model := &SettingModel{ID: s.ID}
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return &setting.Setting{ID: model.ID}, nil
}

func (r *SettingRepository) GetByID(id string) (*setting.Setting, error) {
	var model SettingModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &setting.Setting{ID: model.ID}, nil
}

func (r *SettingRepository) Update(s *setting.Setting) (*setting.Setting, error) {
	model := &SettingModel{ID: s.ID}
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return &setting.Setting{ID: model.ID}, nil
}

func (r *SettingRepository) Delete(id string) error {
	return r.db.Delete(&SettingModel{}, "id = ?", id).Error
}

func (r *SettingRepository) List() ([]*setting.Setting, error) {
	var models []SettingModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	settings := make([]*setting.Setting, len(models))
	for idx, m := range models {
		settings[idx] = &setting.Setting{ID: m.ID}
	}
	return settings, nil
}
