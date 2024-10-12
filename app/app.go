package app

import "github.com/ducconit/gobase/config"

type App struct {
	Config config.Store
}

func NewApp(opts ...Option) *App {
	app := &App{}
	for _, opt := range opts {
		opt(app)
	}
	return app
}
