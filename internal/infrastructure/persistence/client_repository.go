package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/client"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"gorm.io/gorm"
)

type ClientModel struct {
	gorm.Model
	ID         string `gorm:"primaryKey"`
	FirstName  string
	MiddleName string
	LastName   string
	BirthDay   int
	BirthMonth int
	BirthYear  int
	Country    string
	City       string
	Street     string
	Building   string
	Apartment  string
	IsActive   bool
}

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) (client.ClientRepositoryInterface, error) {
	if err := db.AutoMigrate(&ClientModel{}); err != nil {
		return nil, err
	}
	return &ClientRepository{db: db}, nil
}

func (r *ClientRepository) Create(c *client.Client) (*client.Client, error) {
	model := toClientModel(c)
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return toClientDomain(model), nil
}

func (r *ClientRepository) GetByID(id string) (*client.Client, error) {
	var model ClientModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return toClientDomain(&model), nil
}

func (r *ClientRepository) Update(c *client.Client) (*client.Client, error) {
	model := toClientModel(c)
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return toClientDomain(model), nil
}

func (r *ClientRepository) Delete(id string) error {
	return r.db.Delete(&ClientModel{}, "id = ?", id).Error
}

func (r *ClientRepository) List() ([]*client.Client, error) {
	var models []ClientModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	clients := make([]*client.Client, len(models))
	for i, model := range models {
		clients[i] = toClientDomain(&model)
	}
	return clients, nil
}

func toClientDomain(clientModel *ClientModel) *client.Client {
	return &client.Client{
		ID: clientModel.ID,
		Person: shared.Person{
			Name: shared.HumanName{
				FirstName:  clientModel.FirstName,
				MiddleName: clientModel.MiddleName,
				LastName:   clientModel.LastName,
			},
			BirthDate: shared.BirthDate{
				Day:   clientModel.BirthDay,
				Month: clientModel.BirthMonth,
				Year:  clientModel.BirthYear,
			},
			Address: shared.Address{
				Country:   clientModel.Country,
				City:      clientModel.City,
				Street:    clientModel.Street,
				Building:  clientModel.Building,
				Apartment: clientModel.Apartment,
			},
		},
		IsActive: clientModel.IsActive,
	}
}

func toClientModel(client *client.Client) *ClientModel {
	return &ClientModel{
		ID:         client.ID,
		FirstName:  client.Person.Name.FirstName,
		MiddleName: client.Person.Name.MiddleName,
		LastName:   client.Person.Name.LastName,
		BirthDay:   client.Person.BirthDate.Day,
		BirthMonth: client.Person.BirthDate.Month,
		BirthYear:  client.Person.BirthDate.Year,
		Country:    client.Person.Address.Country,
		City:       client.Person.Address.City,
		Street:     client.Person.Address.Street,
		Building:   client.Person.Address.Building,
		Apartment:  client.Person.Address.Apartment,
		IsActive:   client.IsActive,
	}
}
