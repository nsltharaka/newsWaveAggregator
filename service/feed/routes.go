package feed

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nsltharaka/newsWaveAggregator/database"
	"github.com/nsltharaka/newsWaveAggregator/service/auth"
	"github.com/nsltharaka/newsWaveAggregator/types"
	"github.com/nsltharaka/newsWaveAggregator/utils"
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

	r.Post("/create-topic-feeds", h.handleCreateTopicWithFeeds)

	return r
}

func (h *Handler) handleCreateTopicWithFeeds(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(auth.ContextKey("authUser")).(int)

	payload, err := utils.ValidateInput(w, r, &types.CreateFeedPayload{})
	if err != nil {
		return
	}

	performFeedTransaction(r, h.db, int32(userId), payload)

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"userid":  userId,
		"payload": payload,
	})

}
