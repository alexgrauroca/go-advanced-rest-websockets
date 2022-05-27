package middlewares

import (
	"go-advanced-rest-websockets/helpers"
	"go-advanced-rest-websockets/server"
	"log"
	"net/http"
	"strings"
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
			log.Printf("New request: %+v", r)

			if !shouldCheckToken(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}

			_, err := helpers.GetJWTAuthorizationInfo(s, w, r)

			if err != nil {
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
