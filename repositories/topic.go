package repositories

import (
	"context"
	"fmt"

	"github.com/alae-touba/playing-with-go-chi/constants/errs"
	"github.com/alae-touba/playing-with-go-chi/models"
	"github.com/alae-touba/playing-with-go-chi/repositories/ent"
	"github.com/alae-touba/playing-with-go-chi/repositories/ent/topic"
	"github.com/alae-touba/playing-with-go-chi/repositories/ent/user"
	"github.com/google/uuid"
	"go.uber.org/zap"
)


type TopicRepository struct {
	client *ent.Client
	logger *zap.Logger
}

func NewTopicRepository(client *ent.Client, logger *zap.Logger) *TopicRepository {
    return &TopicRepository{
        client: client,
		logger: logger,
    }
}

func (r *TopicRepository) Create(ctx context.Context, req *models.TopicRequest) (*ent.Topic, error) {
    userID, err := uuid.Parse(req.UserID)
    if err != nil {
        return nil, err
    }

    topic, err := r.client.Topic.Create().
        SetName(req.Name).
        SetDescription(req.Description).
        SetImageName(req.ImageName).
        SetUserID(userID).
        Save(ctx)

	if err != nil {
		if ent.IsValidationError(err) {
			return nil, errs.ErrValidationFailed
		}
		return nil, fmt.Errorf("creating topic: %w", err)
	}
	return topic, nil
}

func (r *TopicRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.Topic, error) {	
    topic, err :=  r.client.Topic.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errs.ErrTopicNotFound
		}
		return nil, fmt.Errorf("getting topic: %w", err)
	}
	return topic, nil
}

func (r *TopicRepository) GetTopics(ctx context.Context, limit, offset int, name string, userId uuid.UUID) ([]*ent.Topic, *int, error) {
    query := r.client.Topic.Query()
    
    if name != "" {
        query = query.Where(topic.NameContains(name))
    }

	if userId != uuid.Nil {
		query = query.Where(topic.HasUserWith(user.IDEQ(userId)))
	}

    total, err := query.Count(ctx)
    if err != nil {
        return nil, nil, err
    }

    topics, err := query.
        Limit(limit).
        Offset(offset).
        // Order(ent.Desc(topic.FieldCreatedAt)).
        All(ctx)

	if err != nil {
		return nil, nil, fmt.Errorf("getting topics: %w", err)
	}

    return topics, &total, err
}

func (r *TopicRepository) UpdateTopic(ctx context.Context, id uuid.UUID, req *models.TopicUpdateRequest) (*ent.Topic, error) {
    update := r.client.Topic.UpdateOneID(id)

	if req.Name != "" {
		update.SetName(req.Name)
	}

	if req.Description != "" {
		update.SetDescription(req.Description)
	}

	if req.ImageName != "" {
		update.SetImageName(req.ImageName)
	}
	

	topic, err := update.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errs.ErrTopicNotFound
		}
		return nil, fmt.Errorf("patching topic: %w", err)
	}

	return topic, nil
}

func (r *TopicRepository) Delete(ctx context.Context, id uuid.UUID) error {
    return r.client.Topic.DeleteOneID(id).Exec(ctx)
}
