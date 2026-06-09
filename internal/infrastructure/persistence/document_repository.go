package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/document"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"gorm.io/gorm"
)

type DocumentModel struct {
	gorm.Model
	ID               string `gorm:"primaryKey"`
	UserID           string
	ClientID         string
	DealID           string
	ContractID       string
	CompanyID        string
	Type             string
	Number           string
	IssuedBy         string
	StorageService   string
	StorageURL       string
	IssuedDate       int64
	ExpirationDate   int64
}

type DocumentRepository struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) (document.DocumentRepositoryInterface, error) {
	if err := db.AutoMigrate(&DocumentModel{}); err != nil {
		return nil, err
	}
	return &DocumentRepository{db: db}, nil
}

func (r *DocumentRepository) Create(d *document.Document) (*document.Document, error) {
	model := toDocumentModel(d)
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return toDocumentDomain(model), nil
}

func (r *DocumentRepository) GetByID(id string) (*document.Document, error) {
	var model DocumentModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return toDocumentDomain(&model), nil
}

func (r *DocumentRepository) Update(d *document.Document) (*document.Document, error) {
	model := toDocumentModel(d)
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return toDocumentDomain(model), nil
}

func (r *DocumentRepository) Delete(id string) error {
	return r.db.Delete(&DocumentModel{}, "id = ?", id).Error
}

func (r *DocumentRepository) List() ([]*document.Document, error) {
	var models []DocumentModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	documents := make([]*document.Document, len(models))
	for i, m := range models {
		documents[i] = toDocumentDomain(&m)
	}
	return documents, nil
}

func toDocumentModel(d *document.Document) *DocumentModel {
	return &DocumentModel{
		ID:               d.ID,
		UserID:           d.UserID,
		ClientID:         d.ClientID,
		DealID:           d.DealID,
		ContractID:       d.ContractID,
		CompanyID:        d.CompanyID,
		Type:             d.Type,
		Number:           d.Number,
		IssuedBy:         d.IssuedBy,
		StorageService:   d.StorageReference.ServiceName,
		StorageURL:       d.StorageReference.URL,
		IssuedDate:       d.IssuedDate,
		ExpirationDate:   d.ExpirationDate,
	}
}

func toDocumentDomain(m *DocumentModel) *document.Document {
	return &document.Document{
		ID:               m.ID,
		UserID:           m.UserID,
		ClientID:         m.ClientID,
		DealID:           m.DealID,
		ContractID:       m.ContractID,
		CompanyID:        m.CompanyID,
		Type:             m.Type,
		Number:           m.Number,
		IssuedBy:         m.IssuedBy,
		StorageReference: shared.StorageReference{ServiceName: m.StorageService, URL: m.StorageURL},
		IssuedDate:       m.IssuedDate,
		ExpirationDate:   m.ExpirationDate,
		CreatedAt:        m.CreatedAt.Unix(),
	}
}
