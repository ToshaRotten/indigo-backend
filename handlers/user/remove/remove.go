package remove

import "github.com/gofiber/fiber/v3"

type Remover struct{}

func (Remover) Remove(c fiber.Ctx) error {
	return c.SendString("Delete User")
}
