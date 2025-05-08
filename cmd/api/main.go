package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"my/perfectPetProjectHttp/internal/handlers/http/middleware"
	"my/perfectPetProjectHttp/internal/handlers/http/root"
)

var example = `
curl -X POST "http://localhost:8080/api/users/123/roles?expires_in=7d" \
     -H "Content-Type: application/json" \
     -H "X-Request-ID: abc123" \
     -H "Authorization: ansbdjahsbdkhasj" \
     -d '{"role": "admin"}'
`

func main() {
	mw := middleware.Chain(
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		middleware.LogUserIdMiddleware,
		middleware.RequireJSON,
		middleware.RecoverMiddleware,
	)

	handler := root.New()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/users/{user_id}/roles", mw(handler.Handle))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
