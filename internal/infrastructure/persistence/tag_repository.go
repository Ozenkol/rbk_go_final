package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/tag"
	"gorm.io/gorm"
)

type TagModel struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	ClientID string
	Name     string
}

type TagRepository struct {
	DB *gorm.DB
}

func NewTagRepository(db *gorm.DB) (tag.TagRepositoryInterface, error) {
	if err := db.AutoMigrate(&TagModel{}); err != nil {
		return nil, err
	}
	return &TagRepository{DB: db}, nil
}

func (r *TagRepository) Create(tag *tag.Tag) error {
	model := toTagModel(tag)
	return r.DB.Create(&model).Error
}

func (r *TagRepository) GetByID(id string) (*tag.Tag, error) {
	var model TagModel
	if err := r.DB.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return toTagDomain(&model), nil
}

func (r *TagRepository) Update(tag *tag.Tag) (*tag.Tag, error) {
	var model TagModel
	if err := r.DB.First(&model, "id = ?", tag.ID).Error; err != nil {
		return nil, err
	}
	model.Name = tag.Name
	if err := r.DB.Save(&model).Error; err != nil {
		return nil, err
	}
	return toTagDomain(&model), nil
}


func (r *TagRepository) Delete(id string) error {
	return r.DB.Delete(&TagModel{}, "id = ?", id).Error
}

func toTagModel(tag *tag.Tag) *TagModel {
	return &TagModel{
		ID:       tag.ID,
		Name:     tag.Name,
	}
}

func toTagDomain(model *TagModel) *tag.Tag {
	return &tag.Tag{
		ID:       model.ID,
		Name:     model.Name,
	}
}