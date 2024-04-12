package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nsltharaka/newsWaveAggregator/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (server *APIServer) Run() error {

	router := chi.NewRouter()

	// middleware
	router.Use(middleware.Logger)

	// health check
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("running well !"))
	})

	// user routes
	router.Mount("/users", user.NewHandler().RegisterRoutes())

	return http.ListenAndServe(server.addr, router)

}
