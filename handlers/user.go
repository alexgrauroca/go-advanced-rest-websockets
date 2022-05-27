package handlers

import (
	"encoding/json"
	"go-advanced-rest-websockets/helpers"
	"go-advanced-rest-websockets/models"
	"go-advanced-rest-websockets/repository"
	"go-advanced-rest-websockets/server"
	"net/http"
)

func MeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := helpers.GetJWTAuthorizationInfo(s, w, r)

		if err != nil {
			return
		}

		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			user, err := repository.GetUserById(r.Context(), claims.UserId)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
