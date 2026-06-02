package http_deps

import "github.com/Ozenkol/rbk-go-final/internal/application"

type Dependencies struct {
	App *application.Application
}

func NewDependencies(app *application.Application) Dependencies {
	return Dependencies{
		App: app,
	}
}