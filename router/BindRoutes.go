package router

import (
	"go-advanced-rest-websockets/handlers"
	"go-advanced-rest-websockets/middlewares"
	"go-advanced-rest-websockets/server"
	"net/http"

	"github.com/gorilla/mux"
)

func BindRoutes(s server.Server, r *mux.Router) {
	addMiddlewares(s, r)
	addRoutes(s, r)
}

func addMiddlewares(s server.Server, r *mux.Router) {
	r.Use(middlewares.CheckAuthMiddleware(s))
}

func addRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/posts", handlers.InsertPostHandler(s)).Methods(http.MethodPost)
}
