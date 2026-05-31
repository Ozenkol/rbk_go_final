package application

import (
	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateUser *command.CreateUserHandler
}

type Queries struct {
	GetUserByID *query.FetchUserHandler
}

func NewApplication(
	createUserHandler command.CreateUserHandler,
	getUserByIDHandler query.FetchUserHandler,
) *Application {
	return &Application{
		Commands: Commands{
			CreateUser: &createUserHandler,
		},
		Queries: Queries{
			GetUserByID: &getUserByIDHandler,
		},
	}
}
