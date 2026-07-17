// START OF FILE backend/internal/shared/logger/logger.go
package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// New configures and returns a new self-contained zerolog.Logger instance.
// This function does not mutate any global state.
func New(env string) zerolog.Logger {
	zerolog.TimeFieldFormat = time.RFC3339

	var writer io.Writer

	if env == "development" {
		writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "2006-01-02 15:04:05",
		}
		// Set dynamic local log levels if desired, e.g. using level writers.
	} else {
		writer = os.Stdout
	}

	return zerolog.New(writer).With().Timestamp().Caller().Logger()
}
