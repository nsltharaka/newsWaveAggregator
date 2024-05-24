package post

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	router.Get("/all", h.handleGetAllPosts)

	return router

}

func (h *Handler) handleGetAllPosts(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(auth.ContextKey("authUser")).(int)

	q := r.URL.Query()
	limitStr := q.Get("limit")
	offsetStr := q.Get("offset")

	limit, limitErr := strconv.Atoi(limitStr)
	offset, offsetErr := strconv.Atoi(offsetStr)

	if limitErr != nil || offsetErr != nil {
		utils.WriteError(w, http.StatusBadRequest, errors.New("invalid query params"))
		return
	}

	resultSet, _ := h.db.GetAllTopicsWithLimitAndOffset(r.Context(), database.GetAllTopicsWithLimitAndOffsetParams{
		UserID: int32(userID),
		Limit:  int32(limit),
		Offset: int32(offset),
	})

	// next page link if available
	next := ""
	if len(resultSet) == limit {
		next = fmt.Sprintf("/posts/all?limit=%d&offset=%d", limit, (offset + limit))
	}

	pageInfo := map[string]any{
		"count": len(resultSet),
		"next":  next,
	}

	posts := []types.OutGoingPostPayload{}
	for _, row := range resultSet {
		posts = append(posts, utils.PostToPostPayload(row))
	}

	log.Println("HANDLE_GET_POSTS payload length=", len(posts))
	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"info":  pageInfo,
		"posts": posts,
	})
}
