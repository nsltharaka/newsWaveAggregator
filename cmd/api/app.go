package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nsltharaka/newsWaveAggregator/aggregator"
	"github.com/nsltharaka/newsWaveAggregator/database"
	"github.com/nsltharaka/newsWaveAggregator/service/auth"
	"github.com/nsltharaka/newsWaveAggregator/service/feed"
	"github.com/nsltharaka/newsWaveAggregator/service/followTopicFeed"
	"github.com/nsltharaka/newsWaveAggregator/service/post"
	"github.com/nsltharaka/newsWaveAggregator/service/search"
	"github.com/nsltharaka/newsWaveAggregator/service/topic"
	"github.com/nsltharaka/newsWaveAggregator/service/user"
)

type APIServer struct {
	addr string
	db   *database.Queries
}

func NewAPIServer(addr string, db *database.Queries) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (server *APIServer) Run() error {

	router := chi.NewRouter()
	authHandler := auth.NewHandler(server.db)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("running fine"))
	})

	router.HandleFunc("GET /user-check", authHandler.WithAuthUser(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})).ServeHTTP)

	// user routes
	userHandler := user.NewHandler(server.db)
	router.Mount("/users", userHandler.RegisterRoutes())

	// forgot password routes
	router.Mount("/auth", authHandler.RegisterRoutes())

	// topic routes
	topicHandler := topic.NewHandler(server.db, authHandler)
	router.Mount("/topics", topicHandler.RegisterRoutes())

	// feed routes
	feedHandler := feed.NewHandler(server.db, authHandler)
	router.Mount("/feeds", feedHandler.RegisterRoutes())

	// follow_topic_feed routes
	followTopicFeedHandler := followTopicFeed.NewHandler(server.db, authHandler)
	router.Mount("/follow-topic-feed", followTopicFeedHandler.RegisterRoutes())

	// post routes
	postHandler := post.NewHandler(server.db, authHandler)
	router.Mount("/posts", postHandler.RegisterRoutes())

	// search routes
	searchHandler := search.NewHandler(server.db, authHandler)
	router.Mount("/search", searchHandler.RegisterRoutes())

	// Aggregator start
	go func() {
		fmt.Printf("aggregator started...")
		aggregator.StartAggregation(10, 1*time.Hour)
	}()

	return http.ListenAndServe(server.addr, router)

}
