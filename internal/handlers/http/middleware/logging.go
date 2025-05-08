package middleware

import (
	"fmt"
	"net/http"

	"my/perfectPetProjectHttp/internal/handlers/http/http_const"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s\n", r.Method, r.URL.Path)

		next(w, r)
	}
}

func LogUserIdMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		UserID := r.Context().Value(http_const.UserIdKey)
		fmt.Printf("user_id = %d\n", UserID)

		next(w, r)
	}
}
