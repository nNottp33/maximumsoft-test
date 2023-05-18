package middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/nNottp33/maximumsoft-test/src/configs"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token, errCheckHeaders := checkHeaders(c)
	if errCheckHeaders != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":  fiber.StatusUnauthorized,
			"error": errCheckHeaders.Error(),
		})
	}

	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(configs.JWT_SECRET), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":  fiber.StatusUnauthorized,
			"error": "Unauthorized",
		})
	}

	return c.Next()
}

func checkHeaders(c *fiber.Ctx) (string, error) {
	token := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
	if token == "" {
		return "", errors.New("invalid token")
	}

	return token, nil
}
