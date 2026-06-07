package user

type UserRepositoryInterface interface {
	Create(user *User) (*User, error)
	GetByID(id string) (*User, error)
	Update(user *User) (*User, error)
	Delete(id string) error
	List() ([]*User, error)
	FindByEmail(email string) (*User, error)
	FindByHumanName(name interface{}) (*User, error)
}
