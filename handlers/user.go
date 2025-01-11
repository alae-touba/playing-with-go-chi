package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alae-touba/playing-with-go-chi/constants"
	"github.com/alae-touba/playing-with-go-chi/models"
	"github.com/alae-touba/playing-with-go-chi/repositories/ent"
	"github.com/alae-touba/playing-with-go-chi/services"
	"github.com/alae-touba/playing-with-go-chi/utils"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type UserHandler struct {
	logger      *zap.Logger
	userService *services.UserService
}

func NewUserHandler(logger *zap.Logger, userService *services.UserService) *UserHandler {
	return &UserHandler{
		logger:      logger,
		userService: userService,
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	h.logger.Info("Inside getUsers")
	users, err := h.userService.GetUsers(r.Context(), username)
	if err != nil {
		h.logger.Error("failed to get users", zap.Error(err))
		utils.RespondWithError(w, http.StatusInternalServerError, constants.ErrFailedToGetUsers)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidUserID)
		return
	}

	user, err := h.userService.GetUser(r.Context(), id)
	if user == nil {
		utils.RespondWithError(w, http.StatusNotFound, constants.ErrUserNotFound)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidRequestBody)
		return
	}

	user, err := h.userService.CreateUser(r.Context(), req.Username, req.Password)
	if err != nil {
		h.logger.Error("failed to create user", zap.Error(err))
		utils.RespondWithError(w, http.StatusInternalServerError, constants.ErrFailedToCreateUser)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidUserID)
		return
	}

	err = h.userService.DeleteUser(r.Context(), id)
	if err != nil {
		h.logger.Error("failed to delete user", zap.Error(err))

		if ent.IsNotFound(err) {
			utils.RespondWithError(w, http.StatusNotFound, constants.ErrUserNotFound)
			return
		}

		utils.RespondWithError(w, http.StatusInternalServerError, "failed to delete user")
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, map[string]interface{}{"message": "User deleted successfully"})
}
