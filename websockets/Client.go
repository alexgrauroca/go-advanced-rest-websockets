package websockets

import "github.com/gorilla/websocket"

type Client struct {
	hub      *Hub
	id       string
	socket   *websocket.Conn
	outbound chan []byte
}
