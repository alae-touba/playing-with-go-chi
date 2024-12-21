package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alae-touba/playing-with-go-chi/constants"
	"github.com/alae-touba/playing-with-go-chi/models"
	"github.com/alae-touba/playing-with-go-chi/services"
	"github.com/alae-touba/playing-with-go-chi/utils"
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
	name := r.URL.Query().Get("name")
	h.logger.Info("Inside getUsers")
	users := h.userService.GetUsers(name)
	utils.RespondWithJSON(w, http.StatusOK, users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	user := h.userService.GetUser(id)
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

	user, err := h.userService.CreateUser(req.Name, req.Password)
	if err != nil {
		h.logger.Error("failed to create user", zap.Error(err))
		utils.RespondWithError(w, http.StatusInternalServerError, constants.ErrFailedToCreateUser)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}
