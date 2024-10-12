package server

import (
	"github.com/ducconit/gobase/app"
	"net/http"
)

type AdapterServer interface {
	http.Handler
	Setup() error
	App() *app.App
}
