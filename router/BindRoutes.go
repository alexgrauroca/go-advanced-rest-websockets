package router

import (
	"go-advanced-rest-websockets/server"

	"github.com/gorilla/mux"
)

func BindRoutes(s server.Server, r *mux.Router) {
	addMiddlewares(s, r)
	addRoutes(s, r)
}
