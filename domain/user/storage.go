package user

import (
	"context"

	"main/db/entity"
)

type Storage interface {
	Create(ctx context.Context, user entity.UserModel) error
	Get(ctx context.Context, userID int) (entity.UserModel, error)
	GetByLogin(ctx context.Context, login string) (entity.UserModel, error)
	Delete(ctx context.Context, userID int) error
}
