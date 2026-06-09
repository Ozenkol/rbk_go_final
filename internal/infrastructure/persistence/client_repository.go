package persistence

import (
	"strings"

	"github.com/Ozenkol/rbk-go-final/internal/domain/client"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"gorm.io/gorm"
)

type ClientModel struct {
	gorm.Model
	ID                string `gorm:"primaryKey"`
	Type              string
	Name              string
	Email             string
	Phone             string
	WhatsApp          string
	IdentificationNum string
	Address           string
	Source            string
	Status            string
	ResponsibleID     string
	Tags              string // stored as comma-separated string
	IsActive          bool
	CreatedAtUnix     int64
	LastContactAtUnix int64
	Comment           string
	// Person details (Physical)
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
	CompanyID  string
	UserID     string
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
	tags := []string{}
	if clientModel.Tags != "" {
		tags = strings.Split(clientModel.Tags, ",")
	}
	return &client.Client{
		ID:                clientModel.ID,
		UserID:            clientModel.UserID,
		CompanyID:         clientModel.CompanyID,
		Type:              shared.ClientType(clientModel.Type),
		Name:              clientModel.Name,
		Email:             clientModel.Email,
		Phone:             clientModel.Phone,
		WhatsApp:          clientModel.WhatsApp,
		IdentificationNum: clientModel.IdentificationNum,
		Address:           clientModel.Address,
		Source:            clientModel.Source,
		Status:            clientModel.Status,
		ResponsibleID:     clientModel.ResponsibleID,
		Tags:              tags,
		CreatedAt:         clientModel.CreatedAtUnix,
		LastContactAt:     clientModel.LastContactAtUnix,
		Comment:           clientModel.Comment,
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
		ID:                client.ID,
		UserID:            client.UserID,
		CompanyID:         client.CompanyID,
		Type:              string(client.Type),
		Name:              client.Name,
		Email:             client.Email,
		Phone:             client.Phone,
		WhatsApp:          client.WhatsApp,
		IdentificationNum: client.IdentificationNum,
		Address:           client.Address,
		Source:            client.Source,
		Status:            client.Status,
		ResponsibleID:     client.ResponsibleID,
		Tags:              strings.Join(client.Tags, ","),
		CreatedAtUnix:     client.CreatedAt,
		LastContactAtUnix: client.LastContactAt,
		Comment:           client.Comment,
		FirstName:         client.Person.Name.FirstName,
		MiddleName:        client.Person.Name.MiddleName,
		LastName:          client.Person.Name.LastName,
		BirthDay:          client.Person.BirthDate.Day,
		BirthMonth:        client.Person.BirthDate.Month,
		BirthYear:         client.Person.BirthDate.Year,
		Country:           client.Person.Address.Country,
		City:              client.Person.Address.City,
		Street:            client.Person.Address.Street,
		Building:          client.Person.Address.Building,
		Apartment:         client.Person.Address.Apartment,
		IsActive:          client.IsActive,
	}
}
