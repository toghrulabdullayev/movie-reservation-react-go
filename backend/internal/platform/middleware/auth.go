package middleware

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func Auth() fiber.Handler {
	return func(c fiber.Ctx) error {
		log.Println("Auth middleware executed")
		c.Locals("isAuth", true)
		return c.Next()
	}
}

//! can't add arguments without a closure
// func Auth(c fiber.Ctx) error {
// 	log.Println("Auth middleware executed")
// 	return c.Next()
// }
