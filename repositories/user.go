package repositories

import (
	"context"

	"github.com/alae-touba/playing-with-go-chi/models"
	"github.com/alae-touba/playing-with-go-chi/repositories/ent"
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

func (r *UserRepository) Create(ctx context.Context, req *models.UserRequest) (*ent.User, error) {
	return r.client.User.Create().
		SetFirstName(req.FirstName).
		SetLastName(req.LastName).
		SetEmail(req.Email).
		SetPassword(req.Password).
		SetImageName(req.ImageName).
		Save(ctx)
}
