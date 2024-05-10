package topic

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/nsltharaka/newsWaveAggregator/database"
	"github.com/nsltharaka/newsWaveAggregator/service/auth"
	"github.com/nsltharaka/newsWaveAggregator/types"
	"github.com/nsltharaka/newsWaveAggregator/utils"
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
	router.Get("/{topicId}", h.handleGetTopic)

	return router

}

func (h *Handler) handleGetALlTopicsForUser(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(auth.ContextKey("authUser")).(int)

	topics, _ := h.db.GetAllTopicsForUserWithSourceCount(r.Context(), int32(userID))

	topicPayload := []types.OutgoingTopicPayload{}
	for _, topic := range topics {

		imageUrl := ""
		if topic.ImgUrl.Valid {
			imageUrl = topic.ImgUrl.String
		}

		topicPayload = append(topicPayload, types.OutgoingTopicPayload{
			ID:          topic.ID,
			Name:        topic.Name,
			UpdatedAt:   topic.UpdatedAt,
			ImgUrl:      imageUrl,
			SourceCount: int(topic.FeedCount),
		})
	}

	utils.WriteJSON(w, http.StatusOK, topicPayload)

}

func (h *Handler) handleGetALlTopics(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all topics"))
}

func (h *Handler) handleGetTopic(w http.ResponseWriter, r *http.Request) {
	// get all the details about this topic.
	// may include all the feeds user added under this topic.

	userID := r.Context().Value(auth.ContextKey("authUser")).(int)
	topicId, err := uuid.Parse(chi.URLParam(r, "topicId"))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid topic id"))
		return
	}

	// get the topic
	topic, err := h.db.GetTopic(r.Context(), topicId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("topic cannot be found for given id"))
		return
	}

	// and all the feeds of it
	feeds, _ := h.db.GetFeedsForUserTopic(r.Context(), database.GetFeedsForUserTopicParams{
		TopicID: topicId,
		UserID:  int32(userID),
	})

	resp := map[string]any{
		"topic": types.OutgoingTopicPayload{
			ID:          topic.ID,
			Name:        topic.Name,
			ImgUrl:      topic.ImgUrl.String,
			UpdatedAt:   topic.UpdatedAt,
			SourceCount: len(feeds),
		},

		"feeds": feeds,
	}

	utils.WriteJSON(w, http.StatusOK, resp)
}
