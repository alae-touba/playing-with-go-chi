package mappings

import (
	"github.com/alae-touba/playing-with-go-chi/models"
	"github.com/alae-touba/playing-with-go-chi/repositories/ent"
)



func ToTopicResponses(topics []*ent.Topic) []models.TopicResponse {
	topicResponses := make([]models.TopicResponse, len(topics))
	for i, topic := range topics {
		topicResponses[i] = *ToTopicResponse(topic)
	}
	return topicResponses
}

func ToTopicResponse(topic *ent.Topic) *models.TopicResponse {
	return &models.TopicResponse{
		ID:          topic.ID.String(),
		Name:        topic.Name,
		Description: topic.Description,
		ImageName:   topic.ImageName,
		CreatedAt:   topic.CreatedAt,
		UpdatedAt:   topic.UpdatedAt,
		UserID:      topic.UserID.String(),
	}
}
