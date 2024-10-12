package api

import (
	"github.com/ducconit/gobase/api/context"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	*fiber.App
	api fiber.Router

	v1 fiber.Router
}

func RegisterRouter(app *fiber.App) *Router {
	app.Use(context.RegisterContextFiber)

	api := app.Group("/api")

	routes := &Router{
		App: app,
		api: api,
		v1:  api.Group("v1"),
	}

	routes.system()

	return routes
}
