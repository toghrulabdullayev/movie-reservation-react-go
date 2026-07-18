package middleware

import (
	fiberzerolog "github.com/gofiber/contrib/v3/zerolog"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
)

// Logger returns a Fiber v3 handler that logs incoming HTTP requests using the injected logger.
func Logger(log *zerolog.Logger) fiber.Handler {
	return fiberzerolog.New(fiberzerolog.Config{
		Logger: log,
		Fields: []string{"latency", "status", "method", "url", "error"},
	})
}
