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
	"go.uber.org/zap"
)

type TopicHandler struct {
	logger       *zap.Logger
	topicService *services.TopicService
}

func NewTopicHandler(logger *zap.Logger, topicService *services.TopicService) *TopicHandler {
	return &TopicHandler{
		logger:       logger,
		topicService: topicService,
	}
}

func (topicHandler *TopicHandler) GetTopics(w http.ResponseWriter, r *http.Request) {
    paginationParams := utils.GetPaginationParams(r)
    name := r.URL.Query().Get("name")
    userId := r.URL.Query().Get("user_id")

	parsedUserId, err := utils.ParseUUID(userId)
	if err != nil {
		topicHandler.logger.Debug("invalid uuid format", zap.String("user_id", userId), zap.Error(err))
		utils.RespondWithError(w, http.StatusBadRequest, errs.ErrInvalidUUID.Error())
		return
	}

    topics, total, err := topicHandler.topicService.GetTopics(r.Context(), paginationParams.Limit, paginationParams.Offset, name, parsedUserId)
    if err != nil {
        topicHandler.logger.Error("failed to get topics", zap.Error(err))
        utils.RespondWithError(w, http.StatusInternalServerError, "failed to get topics")
        return
    }

    utils.RespondWithList(w, http.StatusOK, topics, paginationParams.Page, paginationParams.PerPage, *total)
}

func (topicHandler *TopicHandler) GetTopic(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")

    parsedID, err := utils.ParseUUID(id)
    if err != nil {
        topicHandler.logger.Debug("invalid uuid format", zap.String("id", id), zap.Error(err))
        utils.RespondWithError(w, http.StatusBadRequest, errs.ErrInvalidUUID.Error())
        return
    }

    topic, err := topicHandler.topicService.GetTopic(r.Context(), parsedID)
    if err != nil {
        switch {
		case errors.Is(err, errs.ErrTopicNotFound):
			utils.RespondWithError(w, http.StatusNotFound, errs.ErrTopicNotFound.Error())
		default:
			topicHandler.logger.Error("failed to get topic", zap.Error(err))
			utils.RespondWithError(w, http.StatusInternalServerError, "failed to get topic")
		}
		return
    }

    utils.RespondWithJSON(w, http.StatusOK, topic)
}

func (topicHandler *TopicHandler) CreateTopic(w http.ResponseWriter, r *http.Request) {
    var req models.TopicRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidRequestBody)
		return
	}

    topic, err := topicHandler.topicService.CreateTopic(r.Context(), &req)
	if err != nil {
        topicHandler.logger.Error("failed to create topic", zap.Error(err))

		switch {
		case errors.Is(err, errs.ErrValidationFailed):
			utils.RespondWithError(w, http.StatusBadRequest, errs.ErrValidationFailed.Error())
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, "failed to create topic")
		}
		return
	}

    utils.RespondWithJSON(w, http.StatusCreated, topic)
}

func (topicHandler *TopicHandler) UpdateTopic(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")

    parsedID, err := utils.ParseUUID(id)
    if err != nil {
        topicHandler.logger.Debug("invalid uuid format", zap.String("id", id), zap.Error(err))
        utils.RespondWithError(w, http.StatusBadRequest, errs.ErrInvalidUUID.Error())
        return
    }

    var req models.TopicUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidRequestBody)
		return
	}

    topic, err := topicHandler.topicService.UpdateTopic(r.Context(), parsedID, &req)
    if err != nil {
		switch {
		case errors.Is(err, errs.ErrTopicNotFound):
			utils.RespondWithError(w, http.StatusNotFound, errs.ErrTopicNotFound.Error())
		default:
			topicHandler.logger.Error("failed to patch topic", zap.Error(err))
			utils.RespondWithError(w, http.StatusInternalServerError, "failed to patch topic")
		}
		return
    }

    utils.RespondWithJSON(w, http.StatusOK, topic)
}

func (topicHandler *TopicHandler) DeleteTopic(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")

    parsedID, err := utils.ParseUUID(id)
    if err != nil {
        topicHandler.logger.Debug("invalid uuid format", zap.String("id", id), zap.Error(err))
        utils.RespondWithError(w, http.StatusBadRequest, errs.ErrInvalidUUID.Error())
        return
    }

	err = topicHandler.topicService.DeleteTopic(r.Context(), parsedID)
    if err != nil {
		switch {
		case errors.Is(err, errs.ErrTopicNotFound):
			utils.RespondWithError(w, http.StatusNotFound, errs.ErrTopicNotFound.Error())
		default:
			topicHandler.logger.Error("failed to delete topic", zap.Error(err))
			utils.RespondWithError(w, http.StatusInternalServerError, "failed to delete user")
		}
		return
    }

    utils.RespondWithJSON(w, http.StatusNoContent, nil)
}