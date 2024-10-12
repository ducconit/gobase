package app

import (
	"github.com/ducconit/gobase/config"
)

type Option func(app *App)

func WithConfig(config config.Store) Option {
	return func(app *App) {
		app.Config = config
	}
}
