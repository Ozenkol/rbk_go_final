package application

import (
	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	"github.com/Ozenkol/rbk-go-final/internal/application/service"
)

type Application struct {
	Commands Commands
	Queries  Queries
	Services Services
}

type Commands struct {
	CreateUser *command.CreateUserHandler
}

type Queries struct {
	GetUserByID *query.FetchUserHandler
}

type Services struct {
	AuthService *service.AuthService
}

func NewApplication(
	createUserHandler command.CreateUserHandler,
	getUserByIDHandler query.FetchUserHandler,
	authService service.AuthService,
) *Application {
	return &Application{
		Commands: Commands{
			CreateUser: &createUserHandler,
		},
		Queries: Queries{
			GetUserByID: &getUserByIDHandler,
		},
		Services: Services{
			AuthService: &authService,
		},
	}
}
