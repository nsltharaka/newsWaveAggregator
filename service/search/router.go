package search

import (
	"database/sql"
	"errors"
	"fmt"
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
	db   *database.Queries
	auth *auth.Handler
}

func NewHandler(db *database.Queries, auth *auth.Handler) *Handler {
	return &Handler{db, auth}
}

func (h *Handler) RegisterRoutes() http.Handler {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(h.auth.WithAuthUser)

	router.Get("/", h.handleSearch) // http://localhost:3030/search?q=xxxx

	return router
}

func (h *Handler) handleSearch(w http.ResponseWriter, r *http.Request) {

	userId := r.Context().Value(auth.ContextKey("authUser")).(int)
	limit := 20

	// searchKey string
	query := r.URL.Query()
	searchKey := query.Get("q")
	pageStr := query.Get("page")

	if searchKey == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("parameter q cannot be empty"))
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, errors.New("invalid query param 'page'"))
		return
	}

	// search for topics and posts
	topicsForQuery, _ := h.db.GetTopicsLike(r.Context(), database.GetTopicsLikeParams{
		UserID: int32(userId),
		Column2: sql.NullString{
			Valid:  true,
			String: searchKey,
		},
	})

	postsForQuery, _ := h.db.GetPostsLike(r.Context(), database.GetPostsLikeParams{
		UserID: int32(userId),
		Column2: sql.NullString{
			Valid:  true,
			String: searchKey,
		},
		Limit:  int32(limit),
		Offset: int32((page - 1) * limit),
	})

	fmt.Println(searchKey, len(topicsForQuery), len(postsForQuery))

	// send search results
	next := fmt.Sprintf("/search?q=%s&page=%d", searchKey, page+1)
	if len(postsForQuery) == limit {
		next = ""
	}

	topicPayload := []types.OutgoingTopicPayload{}
	for _, topic := range topicsForQuery {

		imageUrl := ""
		if topic.ImgUrl.Valid {
			imageUrl = topic.ImgUrl.String
		}

		topicPayload = append(topicPayload, types.OutgoingTopicPayload{
			ID:        topic.ID,
			Name:      topic.Name,
			UpdatedAt: topic.UpdatedAt,
			ImgUrl:    imageUrl,
		})
	}

	posts := []types.OutGoingPostPayload{}
	for _, row := range postsForQuery {

		post := types.OutGoingPostPayload{
			PostID:      row.PostID,
			Title:       row.Title,
			Description: "",
			Author:      "",
			PubDate:     row.PubDate,
			PostImage:   "",
			Url:         row.Url,
			FeedID:      row.FeedID,
			FeedUrl:     row.FeedUrl,
			Topic:       "",
		}

		if row.Description.Valid {
			post.Description = row.Description.String
		}

		if row.Author.Valid {
			post.Author = row.Author.String
		}

		if row.PostImage.Valid {
			post.PostImage = row.PostImage.String
		}

		posts = append(posts, post)

	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"topics": topicPayload,
		"posts":  posts,
		"next":   next,
	})

}
