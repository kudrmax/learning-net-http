package main

import (
	"net/http"

	"my/perfectPetProjectHttp/internal/handlers/http/root"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/users/{user_id}/roles", root.Handle)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
