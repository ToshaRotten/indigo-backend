package user

import "main/db/entity"

type userEntity struct {
	storage Storage

	model entity.UserModel
}

func (u *userEntity) Model() entity.UserModel {
	return u.model
}
