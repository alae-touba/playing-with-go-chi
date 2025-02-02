package services

import (
	"context"

	"github.com/alae-touba/playing-with-go-chi/mappings"
	"github.com/alae-touba/playing-with-go-chi/models"
	"github.com/alae-touba/playing-with-go-chi/repositories"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type TopicService struct {
	logger           *zap.Logger
	topicRepository *repositories.TopicRepository
}

func NewTopicService(logger *zap.Logger, topicRepository *repositories.TopicRepository) *TopicService {
	return &TopicService{
		logger:           logger,
		topicRepository: topicRepository,
	}
}

func (topicService *TopicService) CreateTopic(ctx context.Context, req *models.TopicRequest) (*models.TopicResponse, error) {
	topicEnt, err := topicService.topicRepository.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return mappings.ToTopicResponse(topicEnt), nil
}

func (topicService *TopicService) GetTopic(ctx context.Context, id uuid.UUID) (*models.TopicResponse, error) {
	topicEnt, err := topicService.topicRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mappings.ToTopicResponse(topicEnt), nil
}

func (topicService *TopicService) GetTopics(ctx context.Context, limit, offset int, name string, userId uuid.UUID) ([]models.TopicResponse, *int, error) {
	topicsEnt, total, err := topicService.topicRepository.GetTopics(ctx, limit, offset, name, userId)
	if err != nil {
		return nil, nil, err
	}

	return mappings.ToTopicResponses(topicsEnt), total, nil
}

func (topicService *TopicService) UpdateTopic(ctx context.Context, id uuid.UUID, req *models.TopicUpdateRequest) (*models.TopicResponse, error) {
	topicEnt, err := topicService.topicRepository.UpdateTopic(ctx, id, req)
	if err != nil {
		return nil, err
	}

	return mappings.ToTopicResponse(topicEnt), nil
}

func (topicService *TopicService) DeleteTopic(ctx context.Context, id uuid.UUID) error {
	return topicService.topicRepository.Delete(ctx, id)
}
