package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	config := b.Config()
	b.router = mux.NewRouter()
	binder(b, b.router)
	log.Println("Starting server on port", config.Port)

	if err := http.ListenAndServe(config.Port, b.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
