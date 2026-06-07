package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/file"
	"gorm.io/gorm"
)

type FileModel struct {
	ID   string `gorm:"primaryKey"`
	Name string
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

func (r *FileRepository) Create(f *file.File) (*file.File, error) {
	model := &FileModel{ID: f.ID, Name: f.Name}
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return &file.File{ID: model.ID, Name: model.Name}, nil
}

func (r *FileRepository) GetByID(id string) (*file.File, error) {
	var model FileModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &file.File{ID: model.ID, Name: model.Name}, nil
}

func (r *FileRepository) Update(f *file.File) (*file.File, error) {
	model := &FileModel{ID: f.ID, Name: f.Name}
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return &file.File{ID: model.ID, Name: model.Name}, nil
}

func (r *FileRepository) Delete(id string) error {
	return r.db.Delete(&FileModel{}, "id = ?", id).Error
}

func (r *FileRepository) List() ([]*file.File, error) {
	var models []FileModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	files := make([]*file.File, len(models))
	for i, m := range models {
		files[i] = &file.File{ID: m.ID, Name: m.Name}
	}
	return files, nil
}
