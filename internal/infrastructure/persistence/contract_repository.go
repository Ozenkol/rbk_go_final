package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/contract"
	"gorm.io/gorm"
)

type ContractModel struct {
	ID        string `gorm:"primaryKey"`
	ClientID  string
	CompanyID string
	Content   string
}

type ContractRepository struct {
	db *gorm.DB
}

func NewContractRepository(db *gorm.DB) (contract.ContractRepositoryInterface, error) {
	if err := db.AutoMigrate(&ContractModel{}); err != nil {
		return nil, err
	}
	return &ContractRepository{db: db}, nil
}

func (r *ContractRepository) Create(c *contract.Contract) (*contract.Contract, error) {
	model := toContractModel(c)
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return toContractDomain(model), nil
}

func (r *ContractRepository) GetByID(id string) (*contract.Contract, error) {
	var model ContractModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return toContractDomain(&model), nil
}

func (r *ContractRepository) Update(c *contract.Contract) (*contract.Contract, error) {
	model := toContractModel(c)
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return toContractDomain(model), nil
}

func (r *ContractRepository) Delete(id string) error {
	return r.db.Delete(&ContractModel{}, "id = ?", id).Error
}

func (r *ContractRepository) List() ([]*contract.Contract, error) {
	var models []ContractModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	contracts := make([]*contract.Contract, len(models))
	for i, m := range models {
		contracts[i] = toContractDomain(&m)
	}
	return contracts, nil
}

func toContractModel(c *contract.Contract) *ContractModel {
	return &ContractModel{
		ID:        c.ID,
		ClientID:  c.ClientID,
		CompanyID: c.CompanyID,
		Content:   c.Content,
	}
}

func toContractDomain(m *ContractModel) *contract.Contract {
	return &contract.Contract{
		ID:        m.ID,
		ClientID:  m.ClientID,
		CompanyID: m.CompanyID,
		Content:   m.Content,
	}
}
