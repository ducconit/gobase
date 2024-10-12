package api

import "github.com/ducconit/gobase/api/handlers/system"

func (r *Router) system() {
	r.api.Get("/ping", system.Ping)
}
