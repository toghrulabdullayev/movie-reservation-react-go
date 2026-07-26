package auth

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/config"
)

type Handler struct {
	service Service
	cfg     *config.Config
}

func NewHandler(service Service, cfg *config.Config) *Handler {
	return &Handler{
		service: service,
		cfg:     cfg,
	}
}

// * Swagger annotations
// @Summary Get current user
// @Description Get current user information
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} fiber.Map
// @Router /auth/me [get]
func (h *Handler) Me(c fiber.Ctx) error {
	userID := c.Locals("userID")
	role := c.Locals("role")

	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	return c.JSON(fiber.Map{
		"user": fiber.Map{
			"id":   userID,
			"role": role,
		},
	})
}

// * Swagger annotations
// @Summary Register a new user
// @Description Register a new user with username, email, and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Register Request"
// @Success 201 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Router /auth/register [post]
func (h *Handler) Register(c fiber.Ctx) error {
	var req RegisterRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request payload",
		})
	}

	user, err := h.service.Register(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user registered successfully",
		"user": UserResponse{
			ID:       user.ID.String(),
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		},
	})
}

// * Swagger annotations
// @Summary Login a user
// @Description Login a user with email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login Request"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Failure 401 {object} fiber.Map
// @Router /auth/login [post]
func (h *Handler) Login(c fiber.Ctx) error {
	var req LoginRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request payload",
		})
	}

	user, accessToken, refreshToken, err := h.service.Login(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	h.setTokenCookies(c, accessToken, refreshToken)

	return c.JSON(fiber.Map{
		"message": "login successful",
		"user": UserResponse{
			ID:       user.ID.String(),
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		},
	})
}

// * Swagger annotations
// @Summary Refresh access token
// @Description Refresh the access token using a valid refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} fiber.Map
// @Failure 401 {object} fiber.Map
// @Router /auth/refresh [post]
func (h *Handler) Refresh(c fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "refresh token missing",
		})
	}

	user, accessToken, newRefreshToken, err := h.service.Refresh(c.Context(), refreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	h.setTokenCookies(c, accessToken, newRefreshToken)

	return c.JSON(fiber.Map{
		"message": "token refreshed successfully",
		"user": UserResponse{
			ID:       user.ID.String(),
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		},
	})
}

// * Swagger annotations
// @Summary Logout a user
// @Description Logout a user by clearing the access and refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} fiber.Map
// @Router /auth/logout [post]
func (h *Handler) Logout(c fiber.Ctx) error {
	h.clearTokenCookies(c)
	return c.JSON(fiber.Map{
		"message": "logout successful",
	})
}

// * Swagger annotations
// @Summary Login with Google
// @Description Login with Google
// @Tags Auth
// @Accept json
// @Produce json
// @Success 302 {object} fiber.Map
// @Router /auth/google/login [get]
func (h *Handler) GoogleLogin(c fiber.Ctx) error {
	url := h.service.GetGoogleAuthURL()
	return c.Redirect().To(url)
}

// * Swagger annotations
// @Summary Google Callback
// @Description Handle Google authentication callback
// @Tags Auth
// @Accept json
// @Produce json
// @Success 302 {object} fiber.Map
// @Router /auth/google/callback [get]
func (h *Handler) GoogleCallback(c fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "auth code missing from google redirection",
		})
	}

	_, accessToken, refreshToken, err := h.service.HandleGoogleCallback(c.Context(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	h.setTokenCookies(c, accessToken, refreshToken)

	return c.Redirect().To(h.cfg.ClientRedirectURL)
}

func (h *Handler) setTokenCookies(c fiber.Ctx, accessToken, refreshToken string) {
	secure := h.cfg.Env != "development"

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
		Secure:   secure,
		SameSite: "Lax",
		Path:     "/",
		Domain:   h.cfg.CookieDomain,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HTTPOnly: true,
		Secure:   secure,
		SameSite: "Lax",
		Path:     "/",
		Domain:   h.cfg.CookieDomain,
	})
}

func (h *Handler) clearTokenCookies(c fiber.Ctx) {
	secure := h.cfg.Env != "development"

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		Secure:   secure,
		SameSite: "Lax",
		Path:     "/",
		Domain:   h.cfg.CookieDomain,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		Secure:   secure,
		SameSite: "Lax",
		Path:     "/",
		Domain:   h.cfg.CookieDomain,
	})
}
