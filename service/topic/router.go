package topic

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nsltharaka/newsWaveAggregator/database"
	"github.com/nsltharaka/newsWaveAggregator/service/auth"
)

type Handler struct {
	db *database.Queries
}

func NewHandler(db *database.Queries) *Handler {
	return &Handler{db}
}

func (h *Handler) RegisterRoutes() http.Handler {

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(auth.WithAuthUser(h.db))

	router.Get("/", h.handleGetALlTopicsForUser)
	router.Get("/all", h.handleGetALlTopics)

	return router

}

func (h *Handler) handleGetALlTopicsForUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all topics for user"))
}

func (h *Handler) handleGetALlTopics(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all topics"))
}
