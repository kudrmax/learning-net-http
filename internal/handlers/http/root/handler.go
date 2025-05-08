package root

import (
	"encoding/json"
	"net/http"
	"strconv"

	"my/perfectPetProjectHttp/internal/handlers/http_errors"
	"my/perfectPetProjectHttp/internal/services/do"
)

// POST /api/users/{user_id}/roles?expires_in=7d
// Content-Type: application/json
// X-Request-ID: abc123
//
// {
//     "role": "admin"
// }

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

type Body struct {
	Role string `json:"role"`
}

type SuccessResponse struct {
	Message string `json:"message,omitempty"`
	Result  string `json:"result"`
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	// валидация

	if r.Method != "POST" {
		http_errors.SendJSONError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("X-Request-ID") == "" {
		http_errors.SendJSONError(w, "X-Request-ID is required", http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http_errors.SendJSONError(w, "Content-Type is required to be application/json", http.StatusUnsupportedMediaType)
		return
	}

	// query params

	expiresIn := r.URL.Query().Get("expires_in")
	if expiresIn == "" {
		http_errors.SendJSONError(w, "expires_in query param is required", http.StatusBadRequest)
		return
	}

	// path params

	userId, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil {
		http_errors.SendJSONError(w, "invalid user_id path param", http.StatusBadRequest)
		return
	}

	// body

	var body Body
	if err = json.NewDecoder(r.Body).Decode(&body); err != nil {
		http_errors.SendJSONError(w, "invalid body", http.StatusBadRequest)
		return
	}

	role := body.Role
	if role == "" {
		http_errors.SendJSONError(w, "role is required", http.StatusBadRequest)
		return
	}

	// logic

	out := do.Do(do.In{
		UserId:    userId,
		Role:      role,
		ExpiresIn: expiresIn,
	})

	// response

	resp := SuccessResponse{
		Message: "OK",
		Result:  out.Result,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
