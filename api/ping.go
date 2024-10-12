package api

import (
	"github.com/ducconit/gobase/api/context"
	"github.com/ducconit/gobase/api/handlers/system"
)

func (r *Router) system() {
	r.api.Get("/ping", context.WithContextFiber(system.Ping))
}
