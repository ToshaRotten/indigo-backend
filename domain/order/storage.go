package order

import (
	"context"

	"main/db/entity"
)

type Storage interface {
	Create(ctx context.Context, order entity.OrderModel) error
	Get(ctx context.Context, orderID int) (entity.OrderModel, error)
	Update(ctx context.Context, order entity.OrderModel) error
}
