package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/aether-winds/luminary-base-server/internal/middleware"
	"github.com/aether-winds/luminary-base-server/internal/route"
)

type server struct {
	Server

	port int
	mux  *http.ServeMux

	middleware []middleware.Middleware
	routes     []route.Route
}

type Server interface {
	RegisterMiddleware(m ...middleware.Middleware) error
	Start() error
	Stop() error
}

func (s *server) RegisterMiddleware(m ...middleware.Middleware) error {
	if len(m) == 0 {
		return errors.New("no middleware provided")
	}

	var newSnapshot = append(s.middleware, m...)

	if len(newSnapshot) == len(s.middleware) {
		return errors.New("failed to register middleware, no new middleware added")
	}

	if len(newSnapshot) < len(s.middleware) {
		return errors.New("failed to register middleware, middleware count decreased")
	}

	s.middleware = newSnapshot
	return nil
}

func (s *server) RegisterRoutes(r ...route.Route) error {
	if len(r) == 0 {
		return errors.New("no routes provided")
	}

	var newSnapshot = append(s.routes, r...)

	if len(newSnapshot) == len(s.routes) {
		return errors.New("failed to register routes, no new routes added")
	}

	if len(newSnapshot) < len(s.routes) {
		return errors.New("failed to register routes, route count decreased")
	}

	s.routes = newSnapshot
	return nil
}

func (s *server) Start() error {
	for _, route := range s.routes {
		s.mux.HandleFunc(fmt.Sprintf("%s %s", route.GetMethod(), route.GetPath()), func(w http.ResponseWriter, r *http.Request) {
			for _, m := range s.middleware {
				if err := m.Handler(r.Context(), w, r); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			if err := route.Handler(r.Context(), w, r); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.mux)
}

func (s *server) Stop() error {
	return nil
}

func CreateServer(port int) Server {
	return &server{
		port: port,
		mux:  http.NewServeMux(),
	}
}
