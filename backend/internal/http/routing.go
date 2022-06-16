package http

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type Routing struct {
	router *mux.Router
}

type Route struct {
	Path      string
	SubRoutes []Route
	Handlers  map[string]http.HandlerFunc
}

func NewRouter(route Route) http.Handler {
	router := mux.NewRouter()
	buildRoute(router, route)
	return router
}

func buildRoute(router *mux.Router, route Route) {
	path := strings.TrimPrefix(route.Path, "//")
	r := router.PathPrefix(path).Subrouter()
	for _, subRoute := range route.SubRoutes {
		buildRoute(r, subRoute)
	}
	for method, handler := range route.Handlers {
		r.Handle("/", handler).Methods(method)
	}
}

type Handler interface {
	Handlers() map[string]http.HandlerFunc
	Path() string
}

type item struct {
	handler Handler
	items   map[string]*item
}

func (i *item) put(subPath []string, handler Handler) {
	if len(subPath) == 0 {
		panic("subPath is empty")
	}
	subItem := i.upsertSubItem(subPath[0])
	if len(subPath) == 1 {
		subItem.handler = handler
		return
	}
	subItem.put(subPath[1:], handler)
}

func (i *item) upsertSubItem(key string) *item {
	subItem, ok := i.items[key]
	if !ok {
		subItem = &item{
			items: map[string]*item{},
		}
		i.items[key] = subItem
	}
	return subItem
}

func (i *item) buildRoute(path string) ([]string, []Route) {
	paths := []string{}
	routes := []Route{}
	for key, item := range i.items {
		p, r := item.buildRoute(path + "/" + key)
		paths = append(paths, p...)
		routes = append(routes, r...)
	}

	if len(routes) > 1 {
		for i2, _ := range routes {
			routes[i2].Path = strings.TrimPrefix(paths[i2], path)
		}
		var h map[string]http.HandlerFunc
		if i.handler != nil {
			h = i.handler.Handlers()
		}
		return []string{path}, []Route{
			{
				Path:      path,
				SubRoutes: routes,
				Handlers:  h,
			},
		}
	}

	if i.handler == nil {
		return paths, routes
	}
	for i2, _ := range routes {
		routes[i2].Path = strings.TrimPrefix(paths[i2], path)
	}
	return []string{path}, []Route{
		{
			Path:      path,
			SubRoutes: routes,
			Handlers:  i.handler.Handlers(),
		},
	}
}

func BuildRouteByHandlers(handlers ...Handler) []Route {
	i := &item{
		items: map[string]*item{},
	}
	for _, handler := range handlers {
		i.put(strings.Split(handler.Path(), "/"), handler)
	}

	_, r := i.buildRoute("/")
	if r == nil {
		panic("no routes")
	}

	return r
}
