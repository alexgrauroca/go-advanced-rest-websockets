package router

import (
	"go-advanced-rest-websockets/server"
	"go-advanced-rest-websockets/websockets"

	"github.com/gorilla/mux"
)

func BindWebsockets(s server.Server, r *mux.Router) {
	hub := websockets.NewHub()

	r.HandleFunc("/ws", hub.HandleWebSocket)
}
