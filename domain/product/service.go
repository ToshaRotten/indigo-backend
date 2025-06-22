package product

import (
	"context"

	"main/db/entity"
)

func (s Service) Get(productID int) (entity.ProductModel, error) {
	product, err := s.storage.Get(context.Background(), productID)
	if err != nil {
		return entity.ProductModel{}, err
	}

	return product, nil
}

func (s Service) GetAll() ([]entity.ProductModel, error) {
	products, err := s.storage.GetAll(context.Background())
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s Service) Create(product entity.ProductModel) error {
	if err := s.storage.Create(context.Background(), product); err != nil {
		return err
	}

	return nil
}
