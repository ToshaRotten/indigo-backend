package product

import (
	"context"

	"main/db/entity"
)

type Storage interface {
	Create(ctx context.Context, product entity.ProductModel) error
	Get(ctx context.Context, productID int) (entity.ProductModel, error)
	Update(ctx context.Context, product entity.ProductModel) error
	Delete(ctx context.Context, productID int) error
}

