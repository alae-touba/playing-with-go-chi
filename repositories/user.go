package repositories

import (
	"context"
	"fmt"

	"github.com/alae-touba/playing-with-go-chi/constants/errs"
	"github.com/alae-touba/playing-with-go-chi/models"
	"github.com/alae-touba/playing-with-go-chi/repositories/ent"
	"github.com/alae-touba/playing-with-go-chi/repositories/ent/user"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UserRepository struct {
	client *ent.Client
	logger *zap.Logger
}

func NewUserRepository(client *ent.Client, logger *zap.Logger) *UserRepository {
	return &UserRepository{
		client: client,
		logger: logger,
	}
}

func (userRepository *UserRepository) Create(ctx context.Context, req *models.UserRequest) (*ent.User, error) {
	user, err := userRepository.client.User.Create().
		SetFirstName(req.FirstName).
		SetLastName(req.LastName).
		SetEmail(req.Email).
		SetPassword(req.Password).
		SetImageName(req.ImageName).
		Save(ctx)

	if err != nil {
		if ent.IsConstraintError(err) {
			userRepository.logger.Error("constraint violation while creating user",
				zap.String("email", req.Email),
				zap.Error(err))
			return nil, errs.ErrEmailExists
		}
		userRepository.logger.Error("failed to create user", zap.Error(err))
		return nil, fmt.Errorf("creating user: %w", err)
	}

	return user, nil
}

func (userRepository *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	user, err := userRepository.client.User.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errs.ErrUserNotFound
		}
		return nil, fmt.Errorf("getting user: %w", err)
	}

	return user, nil
}

func (userRepository *UserRepository) GetByUsername(ctx context.Context, email string) (*ent.User, error) {
	user, err := userRepository.client.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errs.ErrUserNotFound
		}
		return nil, fmt.Errorf("getting user by email: %w", err)
	}

	return user, nil
}
