package update

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

type Updater struct{}

func (Updater) Update(c fiber.Ctx) error {
	fmt.Println("Updating User")
	return nil
}
