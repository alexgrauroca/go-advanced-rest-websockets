package helpers

import (
	"go-advanced-rest-websockets/models"
	"go-advanced-rest-websockets/server"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

func GetJWTAuthorizationInfo(s server.Server, w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JWTSecret), nil
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	return token, err
}
