package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nsltharaka/newsWaveAggregator/database"
	"github.com/nsltharaka/newsWaveAggregator/service/feed"
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

	// middleware
	router.Use(middleware.Logger)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("running fine"))
	})

	// user routes
	userHandler := user.NewHandler(server.db)
	router.Mount("/users", userHandler.RegisterRoutes())

	// feed routes
	feedHandler := feed.NewHandler(server.db)
	router.Mount("/feeds", feedHandler.RegisterRoutes())

	// topic routes
	topicHandler := topic.NewHandler(server.db)
	router.Mount("/topics", topicHandler.RegisterRoutes())

	return http.ListenAndServe(server.addr, router)

}
