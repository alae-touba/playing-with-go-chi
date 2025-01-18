package mappings

import (
	"github.com/alae-touba/playing-with-go-chi/models"
	"github.com/alae-touba/playing-with-go-chi/repositories/ent"
	"time"
)

func ToUserResponses(users []*ent.User) []models.UserResponse {
	userResponses := make([]models.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = *ToUserResponse(user)
	}
	return userResponses
}

func ToUserResponse(user *ent.User) *models.UserResponse {
	var deletedAt *time.Time
	if !user.DeletedAt.IsZero() {
		deletedAt = &user.DeletedAt
	}

	return &models.UserResponse{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		ImageName: user.ImageName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: deletedAt,
	}
}
