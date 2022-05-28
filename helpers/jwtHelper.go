package helpers

import (
	"go-advanced-rest-websockets/models"
	"go-advanced-rest-websockets/repository"
	"go-advanced-rest-websockets/server"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

func GetJWTAuthorizationInfo(s server.Server, w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(s.Config().JWTSecret), nil
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	return token, err
}

func GetUserByJWT(s server.Server, w http.ResponseWriter, r *http.Request) (*models.User, error) {
	token, err := GetJWTAuthorizationInfo(s, w, r)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
		user, err := repository.GetUserById(r.Context(), claims.UserId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil, err
		}

		return user, nil
	}

	http.Error(w, err.Error(), http.StatusInternalServerError)
	return nil, err
}
