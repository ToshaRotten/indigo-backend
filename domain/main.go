package domain

import (
	"main/db/storage"
	"main/db/wrapper"
	"main/domain/user"
)

var User user.Service

func Initialize(DatabaseWrapper *wrapper.DatabaseWrapper) {
	userStorage := storage.NewUserStorage(DatabaseWrapper.Client)

	userService := user.CreateService(userStorage)

	User = userService
}
