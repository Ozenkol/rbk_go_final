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
	CreateTask *command.CreateTaskHandler
	CreateNote *command.CreateNoteHandler
	CreateClient *command.CreateClientHandler
	DeleteTask *command.DeleteTaskHandler
	UpdateTask *command.UpdateTaskHandler
}

type Queries struct {
	GetUserByID *query.FetchUserHandler
	GetTaskByID *query.FetchTaskByIDHandler
}

type Services struct {
	AuthService *service.AuthService
}

func NewApplication(
	createUserHandler command.CreateUserHandler,
	createTaskHandler command.CreateTaskHandler,
	createNoteHandler command.CreateNoteHandler,
	createClientHandler command.CreateClientHandler,
	deleteTaskHandler command.DeleteTaskHandler,
	updateTaskHandler command.UpdateTaskHandler,
	getUserByIDHandler query.FetchUserHandler,
	getTaskByIDHandler query.FetchTaskByIDHandler,
	authService service.AuthService,
) *Application {
	return &Application{
		Commands: Commands{
			CreateUser: &createUserHandler,
			CreateTask: &createTaskHandler,
			CreateNote: &createNoteHandler,
			CreateClient: &createClientHandler,
			DeleteTask: &deleteTaskHandler,
			UpdateTask: &updateTaskHandler,
		},
		Queries: Queries{
			GetUserByID: &getUserByIDHandler,
			GetTaskByID: &getTaskByIDHandler,
		},
		Services: Services{
			AuthService: &authService,
		},
	}
}
