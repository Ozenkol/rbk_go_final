package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/user"

	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"gorm.io/gorm"
)

type UserModel struct {
	ID             string `gorm:"primaryKey"`
	FirstName      string
	MiddleName     string
	LastName       string
	Email          string
	HashedPassword string
	IsVerified     bool
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepositoryInterface {
	if err := db.AutoMigrate(&UserModel{}); err != nil {
		panic(err) // Handle error properly in production
	}
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *user.User) error {
	userModel := toUserModel(user)
	return r.db.Create(userModel).Error
}

func (r *UserRepository) GetByID(id string) (*user.User, error) {
	var model UserModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return toUserDomain(&model), nil
}

func (r *UserRepository) FindByHumanName(humanName shared.HumanName) (*user.User, error) {
	var model UserModel
	if err := r.db.Where("first_name = ? AND middle_name = ? AND last_name = ?", humanName.FirstName, humanName.MiddleName, humanName.LastName).First(&model).Error; err != nil {
		return nil, err
	}
	return toUserDomain(&model), nil
}

func (r *UserRepository) FindByEmail(email string) (*user.User, error) {
	var model UserModel
	if err := r.db.Where("email = ?", email).First(&model).Error; err != nil {
		return nil, err
	}
	return toUserDomain(&model), nil
}

func (r *UserRepository) Update(user *user.User) error {
	userModel := toUserModel(user)
	return r.db.Save(userModel).Error
}

func (r *UserRepository) Delete(id string) error {
	return r.db.Delete(&UserModel{}, id).Error
}

func (r *UserRepository) LogIn(email string, hashedPassword string) (*user.User, error) {
	var model UserModel
	if err := r.db.Where("email = ? AND hashed_password = ?", email, hashedPassword).First(&model).Error; err != nil {
		return nil, err
	}
	return toUserDomain(&model), nil
}

func (r *UserRepository) LogOut(userID string) error {
	// No action needed for stateless JWT authentication
	return nil
}

func (r *UserRepository) Register(humanName shared.HumanName, email string, hashedPassword string) (*user.User, error) {
	newUser, err := user.NewUser(humanName, email, hashedPassword)
	if err != nil {
		return nil, err
	}

	if err := r.Create(&newUser); err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (r *UserRepository) VerifyEmail(userID string, token string) error {
	var model UserModel
	if err := r.db.First(&model, userID).Error; err != nil {
		return err
	}
	model.IsVerified = true
	return r.db.Save(&model).Error
}

func toUserDomain(userModel *UserModel) *user.User {
	return &user.User{
		ID: userModel.ID,
		HumanName: shared.HumanName{
			FirstName:  userModel.FirstName,
			MiddleName: userModel.MiddleName,
			LastName:   userModel.LastName,
		},
		Email:          userModel.Email,
		HashedPassword: userModel.HashedPassword,
		IsVerified:     userModel.IsVerified,
	}
}

func toUserModel(user *user.User) *UserModel {
	return &UserModel{
		ID:             user.ID,
		FirstName:      user.HumanName.FirstName,
		MiddleName:     user.HumanName.MiddleName,
		LastName:       user.HumanName.LastName,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		IsVerified:     user.IsVerified,
	}
}
