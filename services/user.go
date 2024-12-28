package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/alae-touba/playing-with-go-chi/models"
	"github.com/alae-touba/playing-with-go-chi/repositories"
	"github.com/alae-touba/playing-with-go-chi/repositories/ent"
	"github.com/alae-touba/playing-with-go-chi/security"
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

func (s *UserService) GetUsers(ctx context.Context, username string) ([]models.UserResponse, error) {
	var users []*ent.User
	var err error

	if username == "" {
		users, err = s.userRepository.GetAll(ctx)
	} else {
		user, err := s.userRepository.GetByUsername(ctx, username)
		if err == nil {
			users = []*ent.User{user}
		}
	}

	if err != nil {
		s.logger.Error("Failed to get users", zap.Error(err))
		return nil, fmt.Errorf("Failed to get users: %v", err)
	}
	return convertToUserResponses(users), nil
}

func (s *UserService) CreateUser(ctx context.Context, name, password string) (*models.UserResponse, error) {
	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		return nil, err
	}

	userEnt, err := s.userRepository.Create(ctx, name, hashedPassword)

	if err != nil {
		return nil, fmt.Errorf("Failed to create user: %v", err)
	}

	return convertToUserResponse(userEnt), nil
}

func (s *UserService) GetUser(ctx context.Context, id int) (*models.UserResponse, error) {
	user, err := s.userRepository.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to get user by id", zap.Error(err))
		return nil, fmt.Errorf("Failed to get user by id: %v", err)
	}

	return convertToUserResponse(user), nil
}

func (s *UserService) ValidateCredentials(username, password string) bool {

	user, err := s.userRepository.GetByUsername(context.Background(), username)

	if err != nil {
		s.logger.Error("Failed to get user by username", zap.Error(err))
		return false
	}

	return security.VerifyPassword(password, user.Password)
}

// TODO: change file
// Helper function to convert ent.User slice to UserResponse slice
func convertToUserResponses(users []*ent.User) []models.UserResponse {
	userResponses := make([]models.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = models.UserResponse{
			ID:       strconv.Itoa(user.ID),
			Username: user.Username,
		}
	}
	return userResponses
}

// TODO: change file
// helper function to convert ent.User to UserResponse
func convertToUserResponse(user *ent.User) *models.UserResponse {
	return &models.UserResponse{
		ID:       strconv.Itoa(user.ID),
		Username: user.Username,
	}
}
