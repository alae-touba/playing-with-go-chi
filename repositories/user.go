package repositories

import (
	"context"
	"fmt"
	"time"

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
			return nil, errs.ErrEmailExists
		}
		return nil, fmt.Errorf("creating user: %w", err)
	}
	return user, nil

}

func (userRepository *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	user, err := userRepository.client.User.Query().
		Where(
			user.ID(id),
			user.DeletedAtIsNil(),
		).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errs.ErrUserNotFound
		}
		return nil, fmt.Errorf("getting user: %w", err)
	}

	return user, nil
}

func (userRepository *UserRepository) GetByUsername(ctx context.Context, email string) (*ent.User, error) {
	user, err := userRepository.client.User.Query().
		Where(
			user.EmailEQ(email),
			user.DeletedAtIsNil(),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errs.ErrUserNotFound
		}
		return nil, fmt.Errorf("getting user by email: %w", err)
	}

	return user, nil
}

func (userRepository *UserRepository) UpdateUser(ctx context.Context, id uuid.UUID, req *models.UserRequest) (*ent.User, error) {
	update := userRepository.client.User.UpdateOneID(id)

	if req.FirstName != "" {
		update.SetFirstName(req.FirstName)
	}
	if req.LastName != "" {
		update.SetLastName(req.LastName)
	}
	if req.Email != "" {
		update.SetEmail(req.Email)
	}
	if req.Password != "" {
		update.SetPassword(req.Password)
	}
	if req.ImageName != "" {
		update.SetImageName(req.ImageName)
	}

	user, err := update.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errs.ErrUserNotFound
		}
		if ent.IsConstraintError(err) {
			return nil, errs.ErrEmailExists
		}
		return nil, fmt.Errorf("patching user: %w", err)
	}

	return user, nil
}

func (userRepository *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	now := time.Now()
	_, err := userRepository.client.User.
		UpdateOneID(id).
		SetDeletedAt(now).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return errs.ErrUserNotFound
		}
		return fmt.Errorf("deleting user: %w", err)
	}

	return nil
}
