package context

import (
	"github.com/ducconit/gobase/app"
	"github.com/ducconit/gobase/config"
	"github.com/labstack/echo/v4"
)

type EchoContext struct {
	echo.Context

	App *app.App
}

func (e *EchoContext) Config() config.Store {
	return e.App.Config
}

func EchoHandler(h Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*EchoContext)
		return h(ctx)
	}
}
