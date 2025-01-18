package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alae-touba/playing-with-go-chi/constants"
	"github.com/alae-touba/playing-with-go-chi/constants/errs"
	"github.com/alae-touba/playing-with-go-chi/models"
	"github.com/alae-touba/playing-with-go-chi/services"
	"github.com/alae-touba/playing-with-go-chi/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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

func (userHandler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidRequestBody)
		return
	}

	user, err := userHandler.userService.CreateUser(r.Context(), &req)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrEmailExists):
			utils.RespondWithError(w, http.StatusConflict, errs.ErrEmailExists.Error())
		// case erro♂rs.Is(err, errors.ErrInvalidUser):
		//     utils.RespondWithError(w, http.StatusBadRequest, "Invalid user data")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, "failed to create user")
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}

func (userHandler *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		userHandler.logger.Debug("invalid uuid format",
			zap.String("id", id),
			zap.Error(err))
		utils.RespondWithError(w, http.StatusBadRequest, errs.ErrInvalidUUID.Error())
		return
	}

	user, err := userHandler.userService.GetUser(r.Context(), parsedID)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrUserNotFound):
			utils.RespondWithError(w, http.StatusNotFound, errs.ErrUserNotFound.Error())
		default:
			userHandler.logger.Error("failed to get user", zap.Error(err))
			utils.RespondWithError(w, http.StatusInternalServerError, "failed to get user")
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}
