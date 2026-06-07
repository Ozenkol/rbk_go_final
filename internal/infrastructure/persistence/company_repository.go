package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/company"
	"gorm.io/gorm"
)

type CompanyModel struct {
	ID   string `gorm:"primaryKey"`
	Name string
}

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) (company.CompanyRepositoryInterface, error) {
	if err := db.AutoMigrate(&CompanyModel{}); err != nil {
		return nil, err
	}
	return &CompanyRepository{db: db}, nil
}

func (r *CompanyRepository) Create(c *company.Company) (*company.Company, error) {
	model := &CompanyModel{ID: c.ID, Name: c.Name}
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return &company.Company{ID: model.ID, Name: model.Name}, nil
}

func (r *CompanyRepository) GetByID(id string) (*company.Company, error) {
	var model CompanyModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &company.Company{ID: model.ID, Name: model.Name}, nil
}

func (r *CompanyRepository) Update(c *company.Company) (*company.Company, error) {
	model := &CompanyModel{ID: c.ID, Name: c.Name}
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return &company.Company{ID: model.ID, Name: model.Name}, nil
}

func (r *CompanyRepository) Delete(id string) error {
	return r.db.Delete(&CompanyModel{}, "id = ?", id).Error
}

func (r *CompanyRepository) List() ([]*company.Company, error) {
	var models []CompanyModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	companies := make([]*company.Company, len(models))
	for i, m := range models {
		companies[i] = &company.Company{ID: m.ID, Name: m.Name}
	}
	return companies, nil
}
