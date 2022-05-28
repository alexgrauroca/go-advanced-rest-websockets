package middlewares

import (
	"go-advanced-rest-websockets/server"
	"log"
	"net/http"
)

func LogRequestMiddleware(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("New request: %+v", r)
			next.ServeHTTP(w, r)
		})
	}
}
