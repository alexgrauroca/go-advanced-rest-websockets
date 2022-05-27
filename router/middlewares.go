package router

import (
	"go-advanced-rest-websockets/middlewares"
	"go-advanced-rest-websockets/server"

	"github.com/gorilla/mux"
)

func addMiddlewares(s server.Server, r *mux.Router) {
	r.Use(middlewares.CheckAuthMiddleware(s))
}
