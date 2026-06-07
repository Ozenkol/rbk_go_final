package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/setting"
	"gorm.io/gorm"
)

type SettingModel struct {
	ID        string `gorm:"primaryKey"`
	CompanyID string
	Key       string
	Value     string
}

type SettingRepository struct {
	db *gorm.DB
}

func NewSettingRepository(db *gorm.DB) (setting.SettingRepositoryInterface, error) {
	if err := db.AutoMigrate(&SettingModel{}); err != nil {
		panic(err) // Handle error properly in production
	}
	return &SettingRepository{db: db}, nil
}

func (r *SettingRepository) Create(setting *setting.Setting) error {
	return r.db.Create(toSettingModel(setting)).Error
}

func (r *SettingRepository) GetByID(id string) (*setting.Setting, error) {
	var settingModel SettingModel
	if err := r.db.First(&settingModel, id).Error; err != nil {
		return nil, err
	}
	return fromSettingModel(&settingModel), nil
}

func (r *SettingRepository) Update(setting *setting.Setting) error {
	return r.db.Save(toSettingModel(setting)).Error
}

func (r *SettingRepository) Delete(id string) error {
	return r.db.Delete(&SettingModel{}, id).Error
}

func toSettingModel(s *setting.Setting) *SettingModel {
	return &SettingModel{
		ID:        s.ID,
		CompanyID: s.CompanyID,
		Key:       s.Key,
		Value:     s.Value,
	}
}

func fromSettingModel(settingModel *SettingModel) *setting.Setting {
	return &setting.Setting{
		ID:        settingModel.ID,
		CompanyID: settingModel.CompanyID,
		Key:       settingModel.Key,
		Value:     settingModel.Value,
	}
}