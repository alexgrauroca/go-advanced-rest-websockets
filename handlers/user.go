package handlers

import (
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

		helpers.HttpJsonResponse(w, user)
	}
}
