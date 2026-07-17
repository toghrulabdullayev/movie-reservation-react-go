package middleware

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func RBAC(roles ...string) fiber.Handler {
	return func(c fiber.Ctx) error {
		log.Println("RBAC middleware executed:")
		for _, role := range roles {
			log.Println("Allowed role:", role)
		}
		return c.Next()
	}
}
