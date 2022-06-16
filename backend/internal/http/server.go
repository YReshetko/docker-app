package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const defaultPort = "8080"

type Server struct {
	handlers        []Handler
	middlewares     []mux.MiddlewareFunc
	staticResources string
	port            string
}

type ServerOption func(*Server)

func WithHandler(handler Handler) ServerOption {
	return func(s *Server) {
		s.handlers = append(s.handlers, handler)
	}
}

func WithMiddleware(middleware mux.MiddlewareFunc) ServerOption {
	return func(s *Server) {
		s.middlewares = append(s.middlewares, middleware)
	}
}

func WithStaticResource(path string) ServerOption {
	return func(s *Server) {
		s.staticResources = path
	}
}

func WithPort(port string) ServerOption {
	return func(s *Server) {
		s.port = port
	}
}

func NewServer(options ...ServerOption) *Server {
	s := &Server{
		port: defaultPort,
	}
	for _, option := range options {
		option(s)
	}
	return s
}

func (s *Server) Serve() {
	routes := buildRoutes(s.handlers...)
	var static http.Handler
	if s.staticResources != "" {
		static = http.FileServer(http.Dir(s.staticResources))
	}
	router := newMuxRouter(routes, static)
	if len(s.middlewares) > 0 {
		router = middlewares(router, s.middlewares...)
	}

	fmt.Printf("Serving: http://localhost:%s\n", s.port)
	err := http.ListenAndServe(":"+s.port, router)
	if err != nil {
		panic(err)
	}
}
