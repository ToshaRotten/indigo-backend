package create

import (
	"fmt"

	"github.com/gofiber/fiber/v3"

	"main/db/entity"
	"main/domain"
)

type Creator struct{}

func (Creator) Create(c fiber.Ctx) error {
	product := new(entity.ProductModel)
	if err := c.Bind().JSON(product); err != nil {
		return err
	}

	err := domain.Product.Create(*product)
	if err != nil {
		return c.SendString(fmt.Sprintf("Error creating product: %s", err.Error()))
	}

	return nil
}
