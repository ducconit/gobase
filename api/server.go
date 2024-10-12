package api

import (
	"context"
	"errors"
	"github.com/ducconit/gobase/api/server"
	"github.com/ducconit/gobase/app"
	"log"
	"net/http"
	"sync"
)

type Server struct {
	App *app.App

	// startupMutex is mutex to lock server instance access during server configuration and startup. Useful for to get
	// listener address info (on which interface/port was listener bound) without having data races.
	startupMutex sync.RWMutex

	srv *http.Server
}

func NewServer(h server.AdapterServer) (*Server, error) {
	if err := h.Setup(); err != nil {
		return nil, err
	}

	s := &http.Server{
		// default binding address
		Addr:    "127.0.0.1:1297",
		Handler: h,
	}
	srv := &Server{
		App: h.App(),
		srv: s,
	}

	return srv, nil
}

func (s *Server) setup() error {
	addr := s.App.Config.GetString("api.binding")
	if addr != "" {
		s.srv.Addr = addr
	}
	return nil
}

func (s *Server) Listen(addr ...string) error {
	s.startupMutex.Lock()
	if err := s.setup(); err != nil {
		s.startupMutex.Unlock()
		return err
	}
	s.startupMutex.Unlock()
	if len(addr) > 0 && addr[0] != "" {
		s.srv.Addr = addr[0]
	}
	log.Println("Listening on ", s.srv.Addr)

	if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) OnShutdown(f func()) {
	s.srv.RegisterOnShutdown(f)
}

func (s *Server) Shutdown(ctx context.Context) error {
	// TODO: add hooks
	s.startupMutex.Lock()
	defer s.startupMutex.Unlock()
	if err := s.srv.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
