package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v3"

	"main/db/entity"
	"main/domain"
)

type RegistrationRequest struct {
	Login       string `json:"login"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	ThirdName   string `json:"third_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

// TODO: map all fields

func Registration(c fiber.Ctx) error {
	registrationRequest := RegistrationRequest{}

	if err := c.Bind().JSON(&registrationRequest); err != nil {
		return c.SendString(fmt.Sprintf("not registered: %s", err.Error()))
	}

	user := new(entity.UserModel)

	user.Email = registrationRequest.Email
	hash, err := domain.HashPassword(registrationRequest.Password)
	if err != nil {
		return c.SendString(fmt.Sprintf("not registered: %s", err.Error()))
	}

	user.PasswordHash = hash

	if err := domain.User.Create(*user); err != nil {
		return c.SendString(fmt.Sprintf("not registered: %s", err.Error()))
	}

	return c.JSON(user)
}
