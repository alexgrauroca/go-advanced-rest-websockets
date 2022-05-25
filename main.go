package main

import (
	"context"
	"go-advanced-rest-websockets/router"
	"go-advanced-rest-websockets/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal("Error starting new server: ", err)
	}

	s.Start(router.BindRoutes)
}
