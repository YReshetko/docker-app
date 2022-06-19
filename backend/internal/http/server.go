package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/YReshetko/go-annotation/annotations/rest/routing"
)

const defaultPort = "8080"

type Server struct {
	handlers        []routing.Handler
	middlewares     []mux.MiddlewareFunc
	staticResources string
	port            string
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
	router := routing.New(s.handlers...)
	if s.staticResources != "" {
		static := http.FileServer(http.Dir(s.staticResources))
		router.PathPrefix("/").Handler(static).Methods(http.MethodGet)
	}

	if len(s.middlewares) > 0 {
		router = middlewares(router, s.middlewares...)
	}

	fmt.Printf("Serving: http://localhost:%s\n", s.port)
	err := http.ListenAndServe(":"+s.port, router)
	if err != nil {
		panic(err)
	}
}

type ServerOption func(*Server)

func WithHandler(handler routing.Handler) ServerOption {
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

func middlewares(router *mux.Router, middlewareFuncs ...mux.MiddlewareFunc) *mux.Router {
	_ = router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {

		for _, middlewareFunc := range middlewareFuncs {
			route.Handler(middlewareFunc(route.GetHandler()))
		}
		return nil
	})

	_ = router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		p, _ := route.GetPathTemplate()
		m, _ := route.GetMethods()
		h := route.GetHandler()
		fmt.Println(p, m, h)
		return nil
	})

	return router
}
