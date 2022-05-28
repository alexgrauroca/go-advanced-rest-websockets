package websockets

import (
	"encoding/json"
	"log"
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

func (hub *Hub) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Panicln(err)
		http.Error(w, "could not open websocket connection", http.StatusBadRequest)
		return
	}

	client := NewClient(hub, socket)
	hub.registrer <- client

	go client.Write()
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.registrer:
			hub.onConnect(client)
		case client := <-hub.unregister:
			hub.onDisconnect(client)
		}
	}
}

func (hub *Hub) onConnect(client *Client) {
	log.Println("Client connected", client.socket.RemoteAddr())

	hub.mutex.Lock()
	defer hub.mutex.Unlock()

	client.id = client.socket.RemoteAddr().String()
	hub.clients = append(hub.clients, client)
}

func (hub *Hub) onDisconnect(client *Client) {
	log.Println("Client disconnected", client.socket.RemoteAddr())
	client.socket.Close()

	hub.mutex.Lock()
	defer hub.mutex.Unlock()

	clientId := client.id
	i := -1

	for j := 0; j < len(hub.clients); j++ {
		hc := hub.clients[i]

		if hc.id == clientId {
			i = j
			break
		}
	}

	if i > -1 {
		newLen := len(hub.clients) - 1

		copy(hub.clients[i:], hub.clients[i+1:])

		hub.clients[newLen] = nil
		hub.clients = hub.clients[:newLen]
	}
}

func (hub *Hub) Broadcast(message any, ignore *Client) {
	data, _ := json.Marshal(message)

	for i := 0; i < len(hub.clients); i++ {
		client := hub.clients[i]

		if client != ignore {
			client.outbound <- data
		}
	}
}
