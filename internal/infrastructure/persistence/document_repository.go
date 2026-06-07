package persistence

import "gorm.io/gorm"

type DocumentModel struct {
	ID       string `gorm:"primaryKey"`
	ClientID string
	Type     string
	Number   string
	IssuedBy string
	ServiceName string
	URL         string
	IssuedDate           string
	ExpirationDate       string
}

type DocumentRepository struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) (*DocumentRepository, error) {
	if err := db.AutoMigrate(&DocumentModel{}); err != nil {
		panic(err) // Handle error properly in production
	}
	return &DocumentRepository{db: db}, nil
}

func (r *DocumentRepository) Save(document *DocumentModel) (*DocumentModel, error) {
	err := r.db.Save(document).Error
	if err != nil {
		return nil, err
	}
	return document, nil
}

func (r *DocumentRepository) FindByID(id string) (*DocumentModel, error) {
	var model DocumentModel
	if err := r.db.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *DocumentRepository) FindAll() ([]*DocumentModel, error) {
	var models []DocumentModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	documents := make([]*DocumentModel, len(models))
	for i, model := range models {
		documents[i] = &model
	}
	return documents, nil
}

func (r *DocumentRepository) Update(document *DocumentModel) error {
	return r.db.Save(document).Error
}

func (r *DocumentRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&DocumentModel{}).Error
}