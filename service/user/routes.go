package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes() http.Handler {

	r := chi.NewRouter()

	r.Post("/login", h.handleLogin)

	return r
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("handling user login"))
}
