package context

import (
	"github.com/ducconit/gobase/config"
)

type Handler func(Context) error

type Context interface {
	Config() config.Store

	JSON(int, any) error
}
