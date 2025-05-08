package middleware

import (
	"net/http"
)

type MiddlewareType = func(http.HandlerFunc) http.HandlerFunc

func Chain(middlewares ...MiddlewareType) MiddlewareType {
	return func(final http.HandlerFunc) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			final = middlewares[i](final)
		}
		return final
	}
}
