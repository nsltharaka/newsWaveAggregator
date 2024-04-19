package topic

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nsltharaka/newsWaveAggregator/database"
)

type Handler struct {
	db *database.Queries
}

func NewHandler(db *database.Queries) *Handler {
	return &Handler{db}
}

func (h *Handler) RegisterRoutes() http.Handler {

	r := chi.NewRouter()

	r.Get("/all", h.handleGetAllRoutes)

	return r

}

func (h *Handler) handleGetAllRoutes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("handling get all topics"))
}
