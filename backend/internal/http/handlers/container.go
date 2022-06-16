package handlers

import (
	"fmt"
	"net/http"
)

type Container struct{}

func (c *Container) GetContainer(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
	fmt.Fprint(w, `[
							{"title":"Container 1","id":"1", "text":"some text 1", "buttonText": "click me 1"},
							{"title":"Container 2","id":"2", "text":"some text 2", "buttonText": "click me 2"},
							{"title":"Container 3","id":"3", "text":"some text 3", "buttonText": "click me 3"},
							{"title":"Container 4","id":"4", "text":"some text 4", "buttonText": "click me 4"},
							{"title":"Container 5","id":"5", "text":"some text 5", "buttonText": "click me 5"}
					]`)
}

func (c *Container) PatchContainer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `[
							{"title":"Container 1","id":"1", "text":"some text 1", "buttonText": "click me 1"},
							{"title":"Container 2","id":"2", "text":"some text 2", "buttonText": "click me 2"},
							{"title":"Container 3","id":"3", "text":"some text 3", "buttonText": "click me 3"},
							{"title":"Container 4","id":"4", "text":"some text 4", "buttonText": "click me 4"},
							{"title":"Container 5","id":"5", "text":"some text 5", "buttonText": "click me 5"}
					]`)
}

func (c *Container) DeleteContainer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `[
							{"title":"Container 1","id":"1", "text":"some text 1", "buttonText": "click me 1"},
							{"title":"Container 2","id":"2", "text":"some text 2", "buttonText": "click me 2"},
							{"title":"Container 3","id":"3", "text":"some text 3", "buttonText": "click me 3"},
							{"title":"Container 4","id":"4", "text":"some text 4", "buttonText": "click me 4"},
							{"title":"Container 5","id":"5", "text":"some text 5", "buttonText": "click me 5"}
					]`)
}

func (c *Container) Handlers() map[string]http.HandlerFunc {
	return map[string]http.HandlerFunc{
		http.MethodGet:    c.GetContainer,
		http.MethodPatch:  c.PatchContainer,
		http.MethodDelete: c.DeleteContainer,
	}

}

func (c *Container) Path() string {
	return "/api/v1/containers/{container_id}"
}
