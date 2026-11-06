package app

import (
	"context"
)

type Protocol interface {
	SetPublic(b, m int)
	GetBig(small int) int
	CalcZ(big, small int) int
}

// App struct
type App struct {
	p   Protocol
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp(p Protocol) *App {
	app := &App{}
	app.p = p
	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) SetPublic(b, m int) {
	a.p.SetPublic(b, m)
}

func (a *App) GetBig(small int) int {
	return a.p.GetBig(small)
}

func (a *App) CalcZ(big, small int) int {
	return a.p.CalcZ(big, small)
}
