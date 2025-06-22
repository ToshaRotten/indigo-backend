package product

import (
	"main/handlers"
	"main/handlers/product/create"
	"main/handlers/product/get"
	"main/handlers/product/remove"
	"main/handlers/product/update"
)

func NewProductController() *handlers.EntityController {
	return &handlers.EntityController{
		Creator: create.Creator{},
		Getter:  get.Getter{},
		Updater: update.Updater{},
		Remover: remove.Remover{},
	}
}
