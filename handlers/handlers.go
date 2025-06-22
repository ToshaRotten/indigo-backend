package handlers

import "github.com/gofiber/fiber/v3"

type Creator interface {
	Create(c fiber.Ctx) error
}

type Updater interface {
	Update(c fiber.Ctx) error
}

type Remover interface {
	Remove(c fiber.Ctx) error
}

type Getter interface {
	Get(c fiber.Ctx) error
	GetAll(c fiber.Ctx) error
}

type EntityController struct {
	Creator
	Getter
	Updater
	Remover
}
