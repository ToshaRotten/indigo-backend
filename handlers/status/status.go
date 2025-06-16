package status

import (
	"github.com/gofiber/fiber/v3"
)

func Status(c fiber.Ctx) error {
	return c.SendString("Is OK")
}
