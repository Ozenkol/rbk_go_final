package command

import (
	"context"
	"testing"

	"github.com/Ozenkol/rbk-go-final/internal/domain/deal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDealRepository struct {
	mock.Mock
}

func (m *MockDealRepository) Create(d *deal.Deal) (*deal.Deal, error) {
	args := m.Called(d)
	return args.Get(0).(*deal.Deal), args.Error(1)
}

func (m *MockDealRepository) GetByID(id string) (*deal.Deal, error) {
	args := m.Called(id)
	return args.Get(0).(*deal.Deal), args.Error(1)
}

func (m *MockDealRepository) Update(d *deal.Deal) (*deal.Deal, error) {
	args := m.Called(d)
	return args.Get(0).(*deal.Deal), args.Error(1)
}

func (m *MockDealRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockDealRepository) List() ([]*deal.Deal, error) {
	args := m.Called()
	return args.Get(0).([]*deal.Deal), args.Error(1)
}

func TestCreateDealHandler_Handle(t *testing.T) {
	repo := new(MockDealRepository)
	handler := NewCreateDealHandler(repo)

	d := &deal.Deal{
		ID:    "test-id",
		Title: "Test Deal",
	}

	repo.On("Create", d).Return(d, nil)

	res, err := handler.Handle(context.Background(), CreateDealCommand{Deal: d})

	assert.NoError(t, err)
	assert.Equal(t, d, res)
	repo.AssertExpectations(t)
}
