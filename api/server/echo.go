package server

import (
	apiContext "github.com/ducconit/gobase/api/context"
	"github.com/ducconit/gobase/api/handlers/system"
	"github.com/ducconit/gobase/app"
	"github.com/ducconit/gobase/config"
	"github.com/labstack/echo/v4"
)

type EchoServer struct {
	*echo.Echo

	cfg config.Store
	app *app.App
}

func NewEchoServer(app *app.App) AdapterServer {
	return &EchoServer{
		Echo: echo.New(),
		cfg:  app.Config,
		app:  app,
	}
}

func (e *EchoServer) Setup() error {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &apiContext.EchoContext{Context: c, App: e.app}

			return next(cc)
		}
	})
	api := e.Group("/api")

	api.GET("/ping", apiContext.EchoHandler(system.Ping))

	return nil
}

func (e *EchoServer) App() *app.App {
	return e.app
}
