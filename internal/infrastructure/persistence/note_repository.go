package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/note"
	"gorm.io/gorm"
)

type NoteModel struct {
	ID        string `gorm:"primaryKey"`
	UserID    string
	ClientID  string
	CompanyID string
	Content   string
}

type NoteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) (note.NoteRepositoryInterface, error) {
	if err := db.AutoMigrate(&NoteModel{}); err != nil {
		return nil, err
	}
	return &NoteRepository{db: db}, nil
}

func (r *NoteRepository) Create(n *note.Note) (*note.Note, error) {
	model := toNoteModel(n)
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return toNoteDomain(model), nil
}

func (r *NoteRepository) GetByID(id string) (*note.Note, error) {
	var model NoteModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return toNoteDomain(&model), nil
}

func (r *NoteRepository) Update(n *note.Note) (*note.Note, error) {
	model := toNoteModel(n)
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return toNoteDomain(model), nil
}

func (r *NoteRepository) Delete(id string) error {
	return r.db.Delete(&NoteModel{}, "id = ?", id).Error
}

func (r *NoteRepository) List() ([]*note.Note, error) {
	var models []NoteModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	notes := make([]*note.Note, len(models))
	for i, m := range models {
		notes[i] = toNoteDomain(&m)
	}
	return notes, nil
}

func toNoteModel(n *note.Note) *NoteModel {
	return &NoteModel{
		ID:        n.ID,
		UserID:    n.UserID,
		ClientID:  n.ClientID,
		CompanyID: n.CompanyID,
		Content:   n.Content,
	}
}

func toNoteDomain(m *NoteModel) *note.Note {
	return &note.Note{
		ID:        m.ID,
		UserID:    m.UserID,
		ClientID:  m.ClientID,
		CompanyID: m.CompanyID,
		Content:   m.Content,
	}
}
