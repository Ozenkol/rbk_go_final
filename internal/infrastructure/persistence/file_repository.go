package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/file"
	"gorm.io/gorm"
)

type FileModel struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	ClientID string
	Name     string
}

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) (file.FileRepositoryInterface, error) {
	if err := db.AutoMigrate(&FileModel{}); err != nil {
		return nil, err
	}
	return &FileRepository{db: db}, nil
}

func (r *FileRepository) Create(file *file.File) error {
	model := toFileModel(file)
	return r.db.Create(&model).Error
}

func (r *FileRepository) GetByID(id string) (*file.File, error) {
	var model FileModel
	if err := r.db.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return toFileDomain(&model), nil
}

func (r *FileRepository) Update(file *file.File) (*file.File, error) {
	var model FileModel
	if err := r.db.Where("id = ?", file.ID).First(&model).Error; err != nil {
		return nil, err
	}
	model.ClientID = file.ClientID
	model.Name = file.Name
	if err := r.db.Save(&model).Error; err != nil {
		return nil, err
	}
	return toFileDomain(&model), nil
}

func (r *FileRepository) Delete(id string) error {
	return r.db.Delete(&FileModel{}, "id = ?", id).Error
}

func (r *FileRepository) FindAll() ([]*file.File, error) {
	var models []FileModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	files := make([]*file.File, len(models))
	for i, model := range models {
		files[i] = toFileDomain(&model)
	}
	return files, nil
}

func toFileModel(file *file.File) *FileModel {
	return &FileModel{
		ID:       file.ID,
		ClientID: file.ClientID,
		Name:     file.Name,
	}
}

func toFileDomain(model *FileModel) *file.File {
	return &file.File{
		ID:       model.ID,
		ClientID: model.ClientID,
		Name:     model.Name,
	}
}