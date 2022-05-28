package models

type WebSocketMessage struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}
