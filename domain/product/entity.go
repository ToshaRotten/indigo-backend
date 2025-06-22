package product

import "main/db/entity"

type productEntity struct {
	storage Storage

	model entity.ProductModel
}

func (u *productEntity) Model() entity.ProductModel {
	return u.model
}
