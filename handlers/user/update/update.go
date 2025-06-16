package update

import "github.com/gofiber/fiber/v3"

type Updater struct{}

func (Updater) Update(c fiber.Ctx) error {
	return nil
}
