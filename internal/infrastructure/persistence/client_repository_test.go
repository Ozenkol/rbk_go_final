package persistence

import (
	"testing"

	"github.com/Ozenkol/rbk-go-final/internal/domain/client"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newTestDB() (*gorm.DB, error) {
    dsn := "host=localhost user=app password=pass dbname=crm_test port=5433 sslmode=disable"
       return gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Silent),
    })
}

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := newTestDB()
	if err != nil {
		t.Skip("Skipping test: DB connection failed")
	}

	err = db.AutoMigrate(&ClientModel{})
	if err != nil {
		t.Fatal(err)
	}

	return db
}
func TestClientRepository_SaveFindByID(t *testing.T) {
	db := setupTestDB(t)
    if db == nil { return }
	repo, err := NewClientRepository(db)
	if err != nil {
		t.Fatal(err)
	}
	c1 := client.Client{
		ID: "1",
		Person: shared.Person{
			Name: shared.HumanName{
				FirstName:  "John",
				MiddleName: "A.",
				LastName:   "Doe",
			},
			BirthDate: shared.BirthDate{
				Day:   1,
				Month: 1,
				Year:  1990,
			},
			Address: shared.Address{
				Country:   "USA",
				City:      "New York",
				Street:    "5th Avenue",
				Building:  "1",
				Apartment: "1A",
			},
			Phone: shared.PhoneNumber{
				CountryCode: "+1",
				Number:      "1234567890",
			},
		},
		IsActive: true,
	}
	createdUser, err := repo.Create(&c1)
	if err != nil {
		t.Fatal(err)
	}
	println(createdUser)

	found, err := repo.GetByID("1")
	if err != nil {
		t.Fatal(err)
	}
	if found.ID != c1.ID || found.Person.Name.FirstName != c1.Person.Name.FirstName {
		t.Fatalf("Expected to find client with ID %s, got %s", c1.ID, found.ID)
	}
}

func TestClientRepository_FindAll(t *testing.T) {
	db := setupTestDB(t)
    if db == nil { return }
	repo, err := NewClientRepository(db)
	if err != nil {
		t.Fatal(err)
	}
	c1 := client.Client{
		ID: "1",
		Person: shared.Person{
			Name: shared.HumanName{
				FirstName:  "John",
				MiddleName: "A.",
				LastName:   "Doe",
			},
			BirthDate: shared.BirthDate{
				Day:   1,
				Month: 1,
				Year:  1990,
			},
			Address: shared.Address{
				Country:   "USA",
				City:      "New York",
				Street:    "5th Avenue",
				Building:  "1",
				Apartment: "1A",
			},
			Phone: shared.PhoneNumber{
				CountryCode: "+1",
				Number:      "1234567890",
			},
		},
		IsActive: true,
	}
	createdUser, err := repo.Create(&c1)
	println(createdUser)
	if err != nil {
		t.Fatal(err)
	}
	c2 := client.Client{
		ID: "2",
		Person: shared.Person{
			Name: shared.HumanName{
				FirstName:  "Jane",
				MiddleName: "B.",
				LastName:   "Smith",
			},
			BirthDate: shared.BirthDate{
				Day:   15,
				Month: 6,
				Year:  1992,
			},
			Address: shared.Address{
				Country:   "USA",
				City:      "Los Angeles",
				Street:    "Hollywood Boulevard",
				Building:  "100",
				Apartment: "2B",
			},
			Phone: shared.PhoneNumber{
				CountryCode: "+1",
				Number:      "0987654321",
			},
		},
		IsActive: true,
	}
	createdUser, err = repo.Create(&c2)
	if err != nil {
		t.Fatal(err)
	}


	allClients, err := repo.List()
	if err != nil {
		t.Fatal(err)
	}

	if len(allClients) != 2 {
		t.Fatalf("Expected 2 clients, got %d", len(allClients))
	}
}

func TestClientRepository_UpdateDelete(t *testing.T) {
	db := setupTestDB(t)
    if db == nil { return }
	repo, err := NewClientRepository(db)
	if err != nil {
		t.Fatal(err)
	}
	c1 := client.Client{
		ID: "1",
		Person: shared.Person{
			Name: shared.HumanName{
				FirstName:  "John",
				MiddleName: "A.",
				LastName:   "Doe",
			},
			BirthDate: shared.BirthDate{
				Day:   1,
				Month: 1,
				Year:  1990,
			},
			Address: shared.Address{
				Country:   "USA",
				City:      "New York",
				Street:    "5th Avenue",
				Building:  "1",
				Apartment: "1A",
			},
			Phone: shared.PhoneNumber{
				CountryCode: "+1",
				Number:      "1234567890",
			},
		},
		IsActive: true,
	}
	createdUser, err := repo.Create(&c1)
	if err != nil {
		t.Fatal(err)
	}
	println(createdUser)
	c1.IsActive = false
	_, err = repo.Update(&c1)
	if err != nil {
		t.Fatal(err)
	}
	found, err := repo.GetByID("1")
	if err != nil {
		t.Fatal(err)
	}
	if found.IsActive != c1.IsActive {
		t.Fatalf("Expected IsActive to be %v, got %v", c1.IsActive, found.IsActive)
	}
	err = repo.Delete("1")
	if err != nil {
		t.Fatal(err)
	}
	_, err = repo.GetByID("1")
	if err == nil {
		t.Fatal("Expected error when finding deleted client, got nil")
	}
}
