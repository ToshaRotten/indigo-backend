package order

import (
	"context"

	"main/db/entity"
)

func (s Service) Create(order entity.OrderModel) error {
	if err := s.storage.Create(context.Background(), order); err != nil {
		return err
	}

	return nil
}

func (s Service) Get(orderID int) (entity.OrderModel, error) {
	order, err := s.storage.Get(context.Background(), orderID)
	if err != nil {
		return entity.OrderModel{}, err
	}

	return order, nil
}

func (s Service) Delete() {
}

func (s Service) Update() {
}
