package main

import (
	"github.com/YReshetko/docker-app/backend/internal/config"
	"github.com/YReshetko/docker-app/backend/internal/http"
	"github.com/YReshetko/docker-app/backend/internal/http/handlers"
	"github.com/YReshetko/docker-app/backend/internal/http/middlewares"
	"github.com/YReshetko/docker-app/backend/internal/model"
)

const (
	webStatic = "./web"
	resources = "./resources"
)

func main() {
	appConfig, err := config.LoadAppConfig(resources + "/config.json")
	if err != nil {
		panic(err)
	}

	m := model.NewModel(appConfig)

	/*d, _ := json.MarshalIndent(appConfig, "", "  ")
	fmt.Println(string(d))*/

	server := http.NewServer(
		http.WithHandler(handlers.NewServices(m)),
		// TODO remove Dummy handlers
		http.WithHandler(&handlers.Container{}),
		http.WithHandler(&handlers.Service{}),
		http.WithHandler(&handlers.Containers{}),

		// TODO update middlewares
		http.WithMiddleware(middlewares.Logging),
		http.WithStaticResource(webStatic),
		http.WithPort("8181"),
	)
	server.Serve()
}
