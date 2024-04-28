package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r.Use(middleware.Logger)

	r.Post("/login", h.handleLogin)
	r.Post("/register", h.handleRegister)

	return r
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	// payload validation
	payload, err := utils.ValidateInput(w, r, &types.RegisterUserPayload{})
	if err != nil {
		return
	}

	// check if user exists
	_, err = h.userRepo.GetUserByEmail(r.Context(), payload.Email)
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

	// OK respond
	utils.WriteJSON(w, http.StatusCreated, types.OutgoingUserPayload{
		Username: createdUser.Username,
		Email:    createdUser.Email,
		ApiKey:   createdUser.ApiKey,
	})
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

	// payload validation
	payload, err := utils.ValidateInput(w, r, &types.LoginUserPayload{})
	if err != nil {
		return
	}

	// check if the user exists
	// otherwise error
	user, err := h.userRepo.GetUserByEmail(r.Context(), payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf(
			"user not found for the given email",
		))
		return
	}

	// user exists, check password
	if !auth.VerifyPassword(user.Password, payload.Password) {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf(
			"invalid password",
		))
		return
	}

	// OK respond
	utils.WriteJSON(w, http.StatusOK, types.OutgoingUserPayload{
		Username: user.Username,
		Email:    user.Email,
		ApiKey:   user.ApiKey,
	})

}
