package auth

import (
	"github.com/gofiber/fiber/v3"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/middleware"
)

func RegisterRoutes(router fiber.Router, handler *Handler) {
	auth := router.Group("/auth")

	auth.Get("/me", middleware.Auth(handler.cfg.JWTSecret), handler.Me)
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
	auth.Post("/refresh", handler.Refresh)
	auth.Post("/logout", handler.Logout)

	auth.Get("/google/login", handler.GoogleLogin)
	auth.Get("/google/callback", handler.GoogleCallback)
}
