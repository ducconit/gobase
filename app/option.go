package app

import (
	"github.com/ducconit/gobase/config"
	"github.com/ducconit/gobase/db"
)

type Option func(app *App)

func WithConfig(config config.Store) Option {
	return func(app *App) {
		app.Config = config
	}
}

func WithDatabase(db *db.DB) Option {
	return func(app *App) {
		app.DB = db
	}
}
