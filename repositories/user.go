package repositories

import (
	"context"

	"github.com/alae-touba/playing-with-go-chi/repositories/ent"
	"github.com/alae-touba/playing-with-go-chi/repositories/ent/user"
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

func (r *UserRepository) Create(ctx context.Context, username, password string) (*ent.User, error) {
	return r.client.User.Create().
		SetUsername(username).
		SetPassword(password).
		Save(ctx)
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*ent.User, error) {
	return r.client.User.Query().
		Where(user.ID(id)).
		Only(ctx)
}

// get all users
func (r *UserRepository) GetAll(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().
		All(ctx)
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*ent.User, error) {
	return r.client.User.Query().
		Where(user.UsernameEQ(username)).
		Only(ctx)
}

func (r *UserRepository) Update(ctx context.Context, id int, password string) (*ent.User, error) {
	return r.client.User.UpdateOneID(id).
		SetPassword(password).
		Save(ctx)
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	return r.client.User.DeleteOneID(id).
		Exec(ctx)
}
