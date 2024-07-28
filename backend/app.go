package backend

import (
	"context"

	"github.com/issueye/build_magic/backend/controller"
	"github.com/issueye/build_magic/backend/initialize"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	initialize.Initialize()

	return &App{}
}

// startup is called at application startup
func (a *App) Startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a *App) DomReady(ctx context.Context) {
	// Add your action here
	a.ctx = ctx

	controller.GetCode().Ctx = ctx
	controller.GetTemplate().Ctx = ctx
	controller.GetBuildProject().Ctx = ctx
	controller.GetVariable().Ctx = ctx
	controller.GetDocs().Ctx = ctx
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) BeforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) Shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (a *App) GetControllers() []interface{} {
	return []interface{}{
		a,
		controller.GetVariable(),
		controller.GetBuildProject(),
		controller.GetTemplate(),
		controller.GetCode(),
		controller.GetDocs(),
	}
}
