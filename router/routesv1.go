package router

import (
	"go-advanced-rest-websockets/handlers"
	"go-advanced-rest-websockets/server"
	"net/http"

	"github.com/gorilla/mux"
)

func addRoutesV1(s server.Server, r *mux.Router) {
	r.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/posts", handlers.InsertPostHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/posts/{id}", handlers.UpdatePostHandler(s)).Methods(http.MethodPut)
	r.HandleFunc("/posts/{id}", handlers.DeletePostHandler(s)).Methods(http.MethodDelete)
}
