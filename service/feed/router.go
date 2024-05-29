package feed

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mmcdole/gofeed"
	"github.com/nsltharaka/newsWaveAggregator/aggregator"
	"github.com/nsltharaka/newsWaveAggregator/database"
	"github.com/nsltharaka/newsWaveAggregator/service/auth"
	"github.com/nsltharaka/newsWaveAggregator/utils"
)

type Handler struct {
	db         *database.Queries
	auth       *auth.Handler
	feedParser *gofeed.Parser
}

func NewHandler(db *database.Queries, auth *auth.Handler) *Handler {
	return &Handler{db, auth, gofeed.NewParser()}
}

func (h *Handler) RegisterRoutes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(h.auth.WithAuthUser)

	router.Get("/refresh", h.handleRefresh)
	router.Get("/verify", h.verifyFeed)

	return router
}

func (h *Handler) handleRefresh(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(auth.ContextKey("authUser")).(int)

	// get all the feeds user follow
	feeds, _ := h.db.GetAllFeedsForUser(r.Context(), int32(userID))

	wg := &sync.WaitGroup{}
	for _, f := range feeds {
		wg.Add(1)
		go aggregator.ScrapeFeeds(wg, h.db, f)
	}
	wg.Wait()

	utils.WriteJSON(w, http.StatusOK, nil)

}

func (h *Handler) verifyFeed(w http.ResponseWriter, r *http.Request) {

	// a feed url comes and need to verify if it is passable or not

	// get the feed url passed to the endpoint
	feedUrl := r.URL.Query().Get("feed-url")

	// feed url can be empty
	if feedUrl == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("query param 'feed-url' is malicious"))
		return
	}

	// check if it is passable
	response := map[string]string{
		"message": "feed is valid",
		"error":   "",
	}
	_, err := h.feedParser.ParseURL(feedUrl)
	if err != nil {
		response["message"] = ""
		response["error"] = "feed is invalid"
	}

	utils.WriteJSON(w, http.StatusOK, response)

}
