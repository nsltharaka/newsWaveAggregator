package feed

import (
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nsltharaka/newsWaveAggregator/aggregator"
	"github.com/nsltharaka/newsWaveAggregator/database"
	"github.com/nsltharaka/newsWaveAggregator/service/auth"
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

	router.Get("/refresh", h.handleRefresh)

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
