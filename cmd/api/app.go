package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nsltharaka/newsWaveAggregator/database"
	"github.com/nsltharaka/newsWaveAggregator/service/feed"
	"github.com/nsltharaka/newsWaveAggregator/service/followTopicFeed"
	"github.com/nsltharaka/newsWaveAggregator/service/post"
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

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("running fine"))
	})

	router.HandleFunc("/refresh", func(w http.ResponseWriter, r *http.Request) {

	})

	// user routes
	userHandler := user.NewHandler(server.db)
	router.Mount("/users", userHandler.RegisterRoutes())

	// topic routes
	topicHandler := topic.NewHandler(server.db)
	router.Mount("/topics", topicHandler.RegisterRoutes())

	// feed routes
	feedHandler := feed.NewHandler(server.db)
	router.Mount("/feeds", feedHandler.RegisterRoutes())

	// follow_topic_feed routes
	followTopicFeedHandler := followTopicFeed.NewHandler(server.db)
	router.Mount("/follow-topic-feed", followTopicFeedHandler.RegisterRoutes())

	// post routes
	postHandler := post.NewHandler(server.db)
	router.Mount("/posts", postHandler.RegisterRoutes())

	// Aggregator start
	// go func() {
	// 	fmt.Printf("aggregator started...")
	// 	aggregator.StartAggregation(2, 30*time.Minute)
	// }()

	return http.ListenAndServe(server.addr, router)

}
