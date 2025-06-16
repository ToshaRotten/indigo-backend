package user

import (
	"main/handlers"
	"main/handlers/user/create"
	"main/handlers/user/get"
	"main/handlers/user/remove"
	"main/handlers/user/update"
)

func NewUserController() *handlers.EntityController {
	return &handlers.EntityController{
		Creator: create.Creator{},
		Getter:  get.Getter{},
		Updater: update.Updater{},
		Remover: remove.Remover{},
	}
}
