package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v3"

	"main/domain"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func Login(c fiber.Ctx) error {
	var request LoginRequest

	if err := c.Bind().JSON(&request); err != nil {
		return err
	}

	hash, err := domain.HashPassword(request.Password)
	if err != nil {
		return c.SendString(fmt.Sprintf("hash of password is not created %s", err.Error()))
	}

	user, err := domain.User.GetByLogin(request.Login)
	if err != nil {
		return c.SendString(fmt.Sprintf("user not found: %s", err.Error()))
	}

	if hash == user.PasswordHash && request.Login == user.Username {
		// Если учетные данные верны, генерируем JWT
		token, err := domain.CreateToken(user.ID, user.Username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"error": "Не удалось создать токен аутентификации"})
		}

		fmt.Printf("token: %s", token)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
	}
	return c.SendString("login error")
}
