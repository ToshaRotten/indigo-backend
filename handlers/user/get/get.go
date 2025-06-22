package get

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"

	"main/domain"
)

type Getter struct{}

func (Getter) Get(c fiber.Ctx) error {
	userID := c.Params("id", "0")

	id, err := strconv.Atoi(userID)
	if err != nil {
		return c.SendString(fmt.Sprintf("wrong request: %s", err.Error()))
	}
	user, err := domain.User.Get(id)
	if err != nil {
		return c.SendString(fmt.Sprintf("wrong request: %s", err.Error()))
	}

	return c.JSON(user)
}

func (g Getter) GetAll(c fiber.Ctx) error {
	users, err := domain.User.GetAll()
	if err != nil {
		return c.SendString(fmt.Sprintf("wrong request: %s", err.Error()))
	}

	return c.JSON(users)
}
