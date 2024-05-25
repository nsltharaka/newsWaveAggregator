package followTopicFeed

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nsltharaka/newsWaveAggregator/database"
	"github.com/nsltharaka/newsWaveAggregator/lib/topicImages"
	"github.com/nsltharaka/newsWaveAggregator/service/auth"
	"github.com/nsltharaka/newsWaveAggregator/types"
	"github.com/nsltharaka/newsWaveAggregator/utils"
)

type Handler struct {
	db          *database.Queries
	imageFinder *topicImages.ImageFinder
	auth        *auth.Handler
}

func NewHandler(db *database.Queries, auth *auth.Handler) *Handler {

	return &Handler{
		db:          db,
		imageFinder: topicImages.NewImageFinder(topicImages.FromGoogleImages),
		auth:        auth,
	}
}

func (h *Handler) RegisterRoutes() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(h.auth.WithAuthUser)

	r.Post("/create", h.handleFollowTopicFeed)
	r.Get("/all", h.handleGetAllFeedsForUser)

	return r

}

func (h *Handler) handleFollowTopicFeed(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(auth.ContextKey("authUser")).(int)

	payload, err := utils.ValidateInput(w, r, &types.IncomingFollowTopicFeedPayload{})
	if err != nil {
		return
	}

	h.performTransaction(r, payload, userID)

	utils.WriteJSON(w, http.StatusCreated, nil)

	go h.imageFinder.UpdateTopic(h.db, payload.Topic)

}

func (h *Handler) handleGetAllFeedsForUser(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(auth.ContextKey("authUser")).(int)

	data, err := h.db.GetAllFeedsGroupedByTopic(r.Context(), int32(userID))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, data)
}
