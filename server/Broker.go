package server

import (
	"go-advanced-rest-websockets/database"
	"go-advanced-rest-websockets/repository"
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

	repo, err := database.NewPostgresRepository(b.config.DatabaseUrl)

	if err != nil {
		log.Fatal("Error connecting to db: ", err)
	}

	repository.SetRepository(repo)

	log.Println("Starting server on port", config.Port)
	if err := http.ListenAndServe(config.Port, b.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
