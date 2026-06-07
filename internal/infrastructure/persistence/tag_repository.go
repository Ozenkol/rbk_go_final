package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/tag"
	"gorm.io/gorm"
)

type TagModel struct {
	ID   string `gorm:"primaryKey"`
}

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) (tag.TagRepositoryInterface, error) {
	if err := db.AutoMigrate(&TagModel{}); err != nil {
		return nil, err
	}
	return &TagRepository{db: db}, nil
}

func (r *TagRepository) Create(t *tag.Tag) (*tag.Tag, error) {
	model := &TagModel{ID: t.ID}
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return &tag.Tag{ID: model.ID}, nil
}

func (r *TagRepository) GetByID(id string) (*tag.Tag, error) {
	var model TagModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &tag.Tag{ID: model.ID}, nil
}

func (r *TagRepository) Update(t *tag.Tag) (*tag.Tag, error) {
	model := &TagModel{ID: t.ID}
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return &tag.Tag{ID: model.ID}, nil
}

func (r *TagRepository) Delete(id string) error {
	return r.db.Delete(&TagModel{}, "id = ?", id).Error
}

func (r *TagRepository) List() ([]*tag.Tag, error) {
	var models []TagModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	tags := make([]*tag.Tag, len(models))
	for idx, m := range models {
		tags[idx] = &tag.Tag{ID: m.ID}
	}
	return tags, nil
}
