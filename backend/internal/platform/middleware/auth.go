package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/shared/jwtutil"
)

// auth checks for a valid access token in secure cookies
func Auth(jwtSecret string) fiber.Handler {
	return func(c fiber.Ctx) error {
		token := c.Cookies("access_token")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized: missing access token",
			})
		}

		claims, err := jwtutil.ParseToken(token, jwtSecret)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized: invalid or expired access token",
			})
		}

		c.Locals("userID", claims.UserID)
		c.Locals("role", claims.Role)
		c.Locals("isAuth", true)

		return c.Next()
	}
}

//! can't add arguments without a closure
// func Auth(c fiber.Ctx) error {
// 	log.Println("Auth middleware executed")
// 	return c.Next()
// }
