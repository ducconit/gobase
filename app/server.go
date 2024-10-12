package app

import (
	"github.com/ducconit/gobase/api"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Server struct {
	App *App

	router *api.Router
}

func NewServerApi(app *App) *Server {
	f := fiber.New()
	router := api.RegisterRouter(f)

	return &Server{
		App: app,

		router: router,
	}
}

func (s *Server) resolveAddressBinding() string {
	addr := s.App.Config.GetString("api.binding")
	if addr == "" {
		return "127.0.0.1:1297"
	}
	return addr
}

func (s *Server) Run(addr ...string) error {
	listenAddr := s.resolveAddressBinding()
	if len(addr) > 0 && addr[0] != "" {
		listenAddr = addr[0]
	}
	return s.router.Listen(listenAddr)
}

func (s *Server) Shutdown(timeout ...time.Duration) error {
	if len(timeout) == 0 {
		return s.router.Shutdown()
	}
	return s.router.ShutdownWithTimeout(timeout[0])
}
