package command

import (
	"context"
	"testing"

	"github.com/Ozenkol/rbk-go-final/internal/domain/client"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockClientRepository struct {
	mock.Mock
}

func (m *MockClientRepository) Create(c *client.Client) (*client.Client, error) {
	args := m.Called(c)
	return args.Get(0).(*client.Client), args.Error(1)
}

func (m *MockClientRepository) GetByID(id string) (*client.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*client.Client), args.Error(1)
}

func (m *MockClientRepository) Update(c *client.Client) (*client.Client, error) {
	args := m.Called(c)
	return args.Get(0).(*client.Client), args.Error(1)
}

func (m *MockClientRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockClientRepository) List() ([]*client.Client, error) {
	args := m.Called()
	return args.Get(0).([]*client.Client), args.Error(1)
}

func TestCreateClientHandler_Handle(t *testing.T) {
	repo := new(MockClientRepository)
	handler := NewCreateClientHandler(repo)

	cl := &client.Client{
		ID:   "client-1",
		Name: "Test Client",
		Type: shared.ClientTypePerson,
	}

	repo.On("Create", cl).Return(cl, nil)

	res, err := handler.Handle(context.Background(), CreateClientCommand{Client: cl})

	assert.NoError(t, err)
	assert.Equal(t, cl, res)
	repo.AssertExpectations(t)
}
