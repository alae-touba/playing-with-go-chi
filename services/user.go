package services

import (
	"context"
	"fmt"
	"time"

	"github.com/alae-touba/playing-with-go-chi/models"
	"github.com/alae-touba/playing-with-go-chi/repositories"
	"github.com/alae-touba/playing-with-go-chi/repositories/ent"
	"github.com/alae-touba/playing-with-go-chi/security"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UserService struct {
	logger         *zap.Logger
	userRepository *repositories.UserRepository
}

func NewUserService(logger *zap.Logger, userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		logger:         logger,
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *models.UserRequest) (*models.UserResponse, error) {
	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// Create a copy of request with hashed password
	userReq := &models.UserRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
		ImageName: req.ImageName,
	}

	userEnt, err := s.userRepository.Create(ctx, userReq)
	if err != nil {
		return nil, err // Pass through the repository error
	}

	return convertToUserResponse(userEnt), nil
}

func (s *UserService) GetUser(ctx context.Context, id uuid.UUID) (*models.UserResponse, error) {
	user, err := s.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertToUserResponse(user), nil
}

func (s *UserService) ValidateCredentials(username, password string) bool {
	user, err := s.userRepository.GetByUsername(context.Background(), username)
	if err != nil {
		return false
	}

	return security.VerifyPassword(password, user.Password)
}

// Helper function to convert ent.User slice to UserResponse slice
func convertToUserResponses(users []*ent.User) []models.UserResponse {
	userResponses := make([]models.UserResponse, len(users))
	for i, user := range users {
		var deletedAt *time.Time
		if !user.DeletedAt.IsZero() {
			deletedAt = &user.DeletedAt
		}

		userResponses[i] = models.UserResponse{
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
	return userResponses
}

// helper function to convert ent.User to UserResponse
func convertToUserResponse(user *ent.User) *models.UserResponse {
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
