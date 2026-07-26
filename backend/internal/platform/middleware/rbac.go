package middleware

import (
	"github.com/gofiber/fiber/v3"
)

// RBAC verifies that the authenticated user possesses an allowed role
func RBAC(roles ...string) fiber.Handler {
	return func(c fiber.Ctx) error {
		roleRaw := c.Locals("role")
		if roleRaw == nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "forbidden: role configuration missing",
			})
		}

		role, ok := roleRaw.(string)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "forbidden: invalid role configuration format",
			})
		}

		allowed := false
		for _, r := range roles {
			if r == role {
				allowed = true
				break
			}
		}

		if !allowed {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "forbidden: insufficient permissions",
			})
		}

		return c.Next()
	}
}
