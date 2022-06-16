package main

import (
	"github.com/YReshetko/docker-app/backend/internal/http"
	"github.com/YReshetko/docker-app/backend/internal/http/handlers"
	http2 "net/http"
)

func main() {
	r := http.BuildRouteByHandlers(&handlers.Container{}, &handlers.Service{}, &handlers.Containers{})
	router := http.NewRouter(r[0])
	http2.ListenAndServe(":8181", router)
	//fmt.Println(r)

	/*fs := http.FileServer(http.Dir("backend/resources"))
	http.Handle("/", fs)

	http.HandleFunc("/containers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[
							{"title":"Container 1","id":"1", "text":"some text 1", "buttonText": "click me 1"},
							{"title":"Container 2","id":"2", "text":"some text 2", "buttonText": "click me 2"},
							{"title":"Container 3","id":"3", "text":"some text 3", "buttonText": "click me 3"},
							{"title":"Container 4","id":"4", "text":"some text 4", "buttonText": "click me 4"},
							{"title":"Container 5","id":"5", "text":"some text 5", "buttonText": "click me 5"}
					]`)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)*/
}
