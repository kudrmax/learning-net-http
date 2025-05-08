package middleware

import (
	"context"
	"net/http"

	"my/perfectPetProjectHttp/internal/handlers/http/http_const"
	"my/perfectPetProjectHttp/internal/handlers/http_errors"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http_errors.SendJSONError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userID, valid, err := ParseWithClaims(tokenString)
		if err != nil || !valid {
			http_errors.SendJSONError(w, "Invalid token", http.StatusUnauthorized)
		}

		AddValueToContextMiddleware(next, http_const.UserIdKey, userID)(w, r)
	}
}

func AddValueToContextMiddleware(next http.HandlerFunc, key, val any) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, key, val)

		next(w, r.WithContext(ctx))
	}
}

// ParseWithClaims Заглушка симулирующая проверку токена
func ParseWithClaims(tokenString string) (userID int, valid bool, err error) {
	return 123, true, nil
}
