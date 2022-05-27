package handlers

import (
	"go-advanced-rest-websockets/helpers"
	"go-advanced-rest-websockets/server"
	"net/http"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		helpers.HttpJsonResponse(w, HomeResponse{
			Message: "Welcome to Platzi Go",
			Status:  true,
		})
	}
}
