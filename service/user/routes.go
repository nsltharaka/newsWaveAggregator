package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nsltharaka/newsWaveAggregator/database"
	"github.com/nsltharaka/newsWaveAggregator/service/auth"
	"github.com/nsltharaka/newsWaveAggregator/types"
	"github.com/nsltharaka/newsWaveAggregator/utils"
)

type Handler struct {
	userRepo *database.Queries
}

func NewHandler(db *database.Queries) *Handler {
	return &Handler{
		userRepo: db,
	}
}

func (h *Handler) RegisterRoutes() http.Handler {

	r := chi.NewRouter()

	r.Post("/login", h.handleLogin)
	r.Post("/register", withValidation(h.handleRegister))

	return r
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("handling user login"))
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	payload := r.Context().Value(payloadKey).(types.RegisterUserPayload)

	// check if user exists
	_, err := h.userRepo.GetUserByEmail(r.Context(), payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	// create user
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	createdUser, err := h.userRepo.CreateUser(r.Context(), database.CreateUserParams{
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Username:  payload.Username,
		Password:  hashedPassword,
		Email:     payload.Email,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf(""))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, types.UserInfo{
		Username: createdUser.Username,
		Email:    createdUser.Email,
	})
}
