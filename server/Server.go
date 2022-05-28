package server

import (
	"context"
	"errors"
	"go-advanced-rest-websockets/websockets"

	"github.com/gorilla/mux"
)

type Server interface {
	Config() *Config
	Hub() *websockets.Hub
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("Port is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("JWTSecret is required")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("DatabaseUrl is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
		hub:    websockets.NewHub(),
	}

	return broker, nil
}
