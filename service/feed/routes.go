package feed

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nsltharaka/newsWaveAggregator/database"
	"github.com/nsltharaka/newsWaveAggregator/service/auth"
)

type Handler struct {
	db *database.Queries
}

func NewHandler(db *database.Queries) *Handler {
	return &Handler{db: db}
}

func (h *Handler) RegisterRoutes() http.Handler {

	r := chi.NewRouter()

	// middleware
	r.Use(auth.WithAuthUser(h.db))

	r.Post("/create", h.handleCreateFeed)

	return r
}

func (h *Handler) handleCreateFeed(w http.ResponseWriter, r *http.Request) {
	_ = r.Context().Value(auth.ContextKey("authUser")).(int32)
}
