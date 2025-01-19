package services

import (
	"context"

	"github.com/alae-touba/playing-with-go-chi/constants/errs"
	"github.com/alae-touba/playing-with-go-chi/mappings"
	"github.com/alae-touba/playing-with-go-chi/models"
	"github.com/alae-touba/playing-with-go-chi/repositories"
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

func (userService *UserService) CreateUser(ctx context.Context, req *models.UserRequest) (*models.UserResponse, error) {
	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, errs.ErrPasswordHashing
	}

	// Create a copy of request with hashed password
	userReq := &models.UserRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
		ImageName: req.ImageName,
	}

	userEnt, err := userService.userRepository.Create(ctx, userReq)
	if err != nil {
		return nil, err
	}

	return mappings.ToUserResponse(userEnt), nil
}

func (userService *UserService) GetUser(ctx context.Context, id uuid.UUID) (*models.UserResponse, error) {
	user, err := userService.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mappings.ToUserResponse(user), nil
}

func (userService *UserService) UpdateUser(ctx context.Context, id uuid.UUID, req *models.UserRequest) (*models.UserResponse, error) {
	if req.Password != "" {
		hashedPassword, err := security.HashPassword(req.Password)
		if err != nil {
			return nil, errs.ErrPasswordHashing
		}
		req.Password = hashedPassword
	}

	user, err := userService.userRepository.UpdateUser(ctx, id, req)
	if err != nil {
		return nil, err
	}

	return mappings.ToUserResponse(user), nil
}

func (userService *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return userService.userRepository.Delete(ctx, id)
}

func (userService *UserService) ValidateCredentials(username, password string) bool {
	user, err := userService.userRepository.GetByUsername(context.Background(), username)
	if err != nil {
		return false
	}

	return security.VerifyPassword(password, user.Password)
}
