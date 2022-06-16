package main

import (
	"github.com/YReshetko/docker-app/backend/internal/http"
	"github.com/YReshetko/docker-app/backend/internal/http/handlers"
	"github.com/YReshetko/docker-app/backend/internal/http/middlewares"
)

const resources = "./web"

func main() {
	server := http.NewServer(
		http.WithHandler(&handlers.Container{}),
		http.WithHandler(&handlers.Service{}),
		http.WithHandler(&handlers.Containers{}),
		http.WithMiddleware(middlewares.Logging),
		http.WithStaticResource(resources),
		http.WithPort("8181"),
	)
	server.Serve()
}
