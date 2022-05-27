package handlers

import (
	"encoding/json"
	"go-advanced-rest-websockets/helpers"
	"go-advanced-rest-websockets/server"
	"net/http"
)

func MeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := helpers.GetUserByJWT(s, w, r)

		if err != nil {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
