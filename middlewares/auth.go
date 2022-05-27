package middlewares

import (
	"go-advanced-rest-websockets/models"
	"go-advanced-rest-websockets/server"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var (
	NO_AUTH_ROUTES = []string{
		"login",
		"signup",
	}
)

func CheckAuthMiddleware(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !shouldCheckToken(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}

			tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
			_, err := jwt.ParseWithClaims(tokenString, models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTSecret), nil
			})

			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func shouldCheckToken(route string) bool {
	for i := 0; i < len(NO_AUTH_ROUTES); i++ {
		if strings.Contains(route, NO_AUTH_ROUTES[i]) {
			return false
		}
	}

	return true
}
