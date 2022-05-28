package websockets

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Checking valid origins, for this project we accept connections from everywhere
		return true
	},
}

type Hub struct {
	clients    []*Client
	registrer  chan *Client
	unregister chan *Client
	mutex      *sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make([]*Client, 0),
		registrer:  make(chan *Client),
		unregister: make(chan *Client),
		mutex:      &sync.Mutex{},
	}
}
