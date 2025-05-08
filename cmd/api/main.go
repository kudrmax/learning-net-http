package main

import (
	"net/http"

	"my/perfectPetProjectHttp/internal/handlers/http/middleware"
	"my/perfectPetProjectHttp/internal/handlers/http/root"
)

func main() {
	mw := middleware.Chain(middleware.LoggingMiddleware, middleware.RequireJSON, middleware.RecoverMiddleware)

	handler := root.New()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/users/{user_id}/roles", mw(handler.Handle))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
