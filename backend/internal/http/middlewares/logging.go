package middlewares

import (
	"fmt"
	http2 "net/http"
)

func Logging(handler http2.Handler) http2.Handler {
	return http2.HandlerFunc(func(w http2.ResponseWriter, r *http2.Request) {
		fmt.Println("1", r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}
