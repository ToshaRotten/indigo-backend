package get

import (
	"fmt"

	"github.com/gofiber/fiber/v3"

	"main/domain"
)

type Getter struct{}

func (g Getter) Get(c fiber.Ctx) error {
	return nil
}

func (g Getter) GetAll(c fiber.Ctx) error {
	products, err := domain.Product.GetAll()
	if err != nil {
		return c.SendString(fmt.Sprintf("wrong request: %s", err.Error()))
	}

	return c.JSON(products)
}
