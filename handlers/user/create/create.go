package create

import (
	"fmt"

	"github.com/gofiber/fiber/v3"

	"main/db/entity"
	"main/domain"
)

type Creator struct{}

func (Creator) Create(c fiber.Ctx) error {
	user := new(entity.UserModel)

	if err := c.Bind().JSON(user); err != nil {
		return err
	}

	if err := domain.User.Create(*user); err != nil {
		return c.SendString(fmt.Sprintf("user is not created %s", err.Error()))
	}

	return c.JSON(user)
}
