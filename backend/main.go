package main

import (
	"fmt"
	"github.com/YReshetko/docker-app/backend/internal/http"
	"github.com/YReshetko/docker-app/backend/internal/http/handlers"
	http2 "net/http"
)

func main() {
	r := http.BuildRouteByHandlers(&handlers.Container{}, &handlers.Service{}, &handlers.Containers{})
	router := http.NewRouter(r[0], http2.FileServer(http2.Dir("resources")))
	http.Middlewares(router, func(handler http2.Handler) http2.Handler {
		return http2.HandlerFunc(func(w http2.ResponseWriter, r *http2.Request) {
			fmt.Println("1", r.URL.Path)
			handler.ServeHTTP(w, r)
		})
	}, func(handler http2.Handler) http2.Handler {
		return http2.HandlerFunc(func(w http2.ResponseWriter, r *http2.Request) {
			fmt.Println("2", r.URL.Path)
			handler.ServeHTTP(w, r)
		})
	})

	fmt.Println("Serving: http://localhost:8181")
	http2.ListenAndServe(":8181", router)
}
