package router

import (
	"go-advanced-rest-websockets/server"

	"github.com/gorilla/mux"
)

func BindWebsockets(s server.Server, r *mux.Router) {
	r.HandleFunc("/ws", s.Hub().HandleWebSocket)
}
