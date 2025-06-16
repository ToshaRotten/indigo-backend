package user

import (
	"context"

	"main/db/entity"
)

func (s Service) Create(user entity.UserModel) error {
	if err := s.storage.Create(context.Background(), user); err != nil {
		return err
	}

	return nil
}

func (s Service) Get(userID int) (entity.UserModel, error) {
	user, err := s.storage.Get(context.Background(), userID)
	if err != nil {
		return entity.UserModel{}, err
	}

	return user, nil
}

func (s Service) Delete() {
}

func (s Service) Update() {
}
