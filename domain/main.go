package domain

import (
	"main/db/storage"
	"main/db/wrapper"
	"main/domain/product"
	"main/domain/user"
)

var (
	User    user.Service
	Product product.Service
)

func Initialize(DatabaseWrapper *wrapper.DatabaseWrapper) {
	userStorage := storage.NewUserStorage(DatabaseWrapper.Client)
	productStorage := storage.NewProductStorage(DatabaseWrapper.Client)

	userService := user.CreateService(userStorage)
	productService := product.CreateService(productStorage)

	User = userService
	Product = productService
}
