package user

import (
	"context"

	"main/db/entity"
)

type Storage interface {
	Create(ctx context.Context, user entity.UserModel) error
	Get(ctx context.Context, userID int) (entity.UserModel, error)
}
