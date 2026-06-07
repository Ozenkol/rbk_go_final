package persistence

import (
	"fmt"

	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"github.com/Ozenkol/rbk-go-final/internal/domain/user"
	"gorm.io/gorm"
)

type UserModel struct {
	ID             string `gorm:"primaryKey"`
	FirstName      string
	LastName       string
	Email          string `gorm:"uniqueIndex"`
	HashedPassword string
	IsVerified     bool
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepositoryInterface {
	if err := db.AutoMigrate(&UserModel{}); err != nil {
		fmt.Printf("Error migrating UserModel: %v\n", err)
	}
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(u *user.User) (*user.User, error) {
	model := toUserModel(u)
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return toUserDomain(model), nil
}

func (r *UserRepository) GetByID(id string) (*user.User, error) {
	var model UserModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return toUserDomain(&model), nil
}

func (r *UserRepository) Update(u *user.User) (*user.User, error) {
	model := toUserModel(u)
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return toUserDomain(model), nil
}

func (r *UserRepository) Delete(id string) error {
	return r.db.Delete(&UserModel{}, "id = ?", id).Error
}

func (r *UserRepository) List() ([]*user.User, error) {
	var models []UserModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	users := make([]*user.User, len(models))
	for i, m := range models {
		users[i] = toUserDomain(&m)
	}
	return users, nil
}

func (r *UserRepository) FindByEmail(email string) (*user.User, error) {
	var model UserModel
	if err := r.db.Where("email = ?", email).First(&model).Error; err != nil {
		return nil, err
	}
	return toUserDomain(&model), nil
}

func (r *UserRepository) FindByHumanName(name interface{}) (*user.User, error) {
	hn, ok := name.(shared.HumanName)
	if !ok {
		return nil, nil
	}
	var model UserModel
	if err := r.db.Where("first_name = ? AND last_name = ?", hn.FirstName, hn.LastName).First(&model).Error; err != nil {
		return nil, err
	}
	return toUserDomain(&model), nil
}

func toUserModel(u *user.User) *UserModel {
	return &UserModel{
		ID:             u.ID,
		FirstName:      u.HumanName.FirstName,
		LastName:       u.HumanName.LastName,
		Email:          u.Email,
		HashedPassword: u.HashedPassword,
		IsVerified:     u.IsVerified,
	}
}

func toUserDomain(m *UserModel) *user.User {
	return &user.User{
		ID: m.ID,
		HumanName: shared.HumanName{
			FirstName: m.FirstName,
			LastName:  m.LastName,
		},
		Email:          m.Email,
		HashedPassword: m.HashedPassword,
		IsVerified:     m.IsVerified,
	}
}
