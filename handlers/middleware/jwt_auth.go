package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"

	"main/domain"
)

func JwtMiddleware(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"error": "Токен аутентификации отсутствует"})
	}

	tokenString := ""
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString = authHeader[7:]
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Неверный формат заголовка Authorization"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(domain.JwtSecret), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{"error": "Неверная подпись токена"})
		}
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"error": "Недействительный токен: " + err.Error()})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Недействительный токен"})
	}

	c.Locals("user", token)
	return c.Next()
}
