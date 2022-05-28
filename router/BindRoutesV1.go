package router

import (
	"go-advanced-rest-websockets/server"

	"github.com/gorilla/mux"
)

func BindRoutesV1(s server.Server, r *mux.Router) {
	api := r.PathPrefix("/api/v1").Subrouter()

	addMiddlewaresV1(s, api)
	addRoutesV1(s, api)
}
