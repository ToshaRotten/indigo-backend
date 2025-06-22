package get

import "github.com/gofiber/fiber/v3"

type Getter struct{}

func (g Getter) Get(c fiber.Ctx) error {
	return nil
}
