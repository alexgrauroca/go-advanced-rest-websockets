package router

import (
	"go-advanced-rest-websockets/handlers"
	"go-advanced-rest-websockets/server"
	"net/http"

	"github.com/gorilla/mux"
)

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
}
