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

type userData struct {
	ID         int    `json:"id"`
	FullName   string `json:"full_name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	CategoryID int    `json:"category_id"`
}

type LoginResponse struct {
	User  userData `json:"userData"`
	Token string   `json:"token"`
}

func Login(c fiber.Ctx) error {
	var request LoginRequest

	if err := c.Bind().JSON(&request); err != nil {
		return err
	}

	user, err := domain.User.GetByLogin(request.Login)
	if err != nil {
		return c.SendString(fmt.Sprintf("user not found: %s", err.Error()))
	}

	if domain.CheckPasswordHash(request.Password, user.PasswordHash) &&
		request.Login == user.Username {
		// Если учетные данные верны, генерируем JWT
		token, err := domain.CreateToken(user.ID, user.Username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"error": "Не удалось создать токен аутентификации"})
		}

		fmt.Printf("token: %s", token)

		response := LoginResponse{
			User: userData{
				ID:         user.ID,
				FullName:   user.FullName,
				Email:      user.Email,
				Username:   user.Username,
				CategoryID: user.UserCategoryID,
			},
			Token: token,
		}

		return c.Status(fiber.StatusOK).JSON(response)
	}
	return c.SendString("login error")
}

