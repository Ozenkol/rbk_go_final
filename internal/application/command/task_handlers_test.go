package command

import (
	"context"
	"testing"

	"github.com/Ozenkol/rbk-go-final/internal/domain/task"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Create(t *task.Task) (*task.Task, error) {
	args := m.Called(t)
	return args.Get(0).(*task.Task), args.Error(1)
}

func (m *MockTaskRepository) GetByID(id string) (*task.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*task.Task), args.Error(1)
}

func (m *MockTaskRepository) Update(t *task.Task) (*task.Task, error) {
	args := m.Called(t)
	return args.Get(0).(*task.Task), args.Error(1)
}

func (m *MockTaskRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskRepository) List() ([]*task.Task, error) {
	args := m.Called()
	return args.Get(0).([]*task.Task), args.Error(1)
}

func TestCreateTaskHandler_Handle(t *testing.T) {
	repo := new(MockTaskRepository)
	handler := NewCreateTaskHandler(repo)

	tk := &task.Task{
		ID:    "task-1",
		Title: "Test Task",
	}

	repo.On("Create", tk).Return(tk, nil)

	res, err := handler.Handle(context.Background(), CreateTaskCommand{Task: tk})

	assert.NoError(t, err)
	assert.Equal(t, tk, res)
	repo.AssertExpectations(t)
}
