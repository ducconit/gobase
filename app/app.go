package app

import (
	"github.com/ducconit/gobase/config"
	"github.com/ducconit/gobase/db"
)

type App struct {
	Config config.Store
	DB     *db.DB
}

func NewApp(opts ...Option) *App {
	app := &App{}
	for _, opt := range opts {
		opt(app)
	}
	return app
}
