In Go, the idiomatic way to manage Dependency Injection (DI) is **constructor injection using interfaces**.

By designing your code to **accept interfaces and return concrete types (or interfaces where appropriate)**, you decouple your business logic from specific databases, frameworks, or external APIs. This also makes unit testing with mocks straightforward.

Below is a complete, structural guide demonstrating how dependencies are passed and wired up across every single layer of your application architecture.

---

### 1. Repository Layer (Database & Cache DI)

Repositories interact directly with databases or caches. They accept concrete database connection pools (like `*pgxpool.Pool` or `*redis.Client`) but are returned as an **interface** so the service layer is not tightly coupled to a database implementation.

#### Example: `movies/repository.go`

```go
package movies

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

// Repository interface defines the data store operations.
// The service layer will accept this interface.
type Repository interface {
	CreateMovie(ctx context.Context, movie *Movie) error
	GetMovieByID(ctx context.Context, id uuid.UUID) (*Movie, error)
}

// repository is the unexported, concrete implementation.
type repository struct {
	db  *pgxpool.Pool
	rdb *redis.Client // We can also inject Redis if the repository handles caching directly
}

// NewRepository is the constructor injecting the raw database connections.
func NewRepository(db *pgxpool.Pool, rdb *redis.Client) Repository {
	return &repository{
		db:  db,
		rdb: rdb,
	}
}

func (r *repository) CreateMovie(ctx context.Context, movie *Movie) error {
	// Use r.db or r.rdb to execute database transactions or cache updates
	return nil
}

func (r *repository) GetMovieByID(ctx context.Context, id uuid.UUID) (*Movie, error) {
	return nil, nil
}
```

---

### 2. Service Layer (Business Logic DI)

Services coordinate business logic. They accept repositories (via **interfaces**) and external utilities (like token generators, hashers, or other services) also defined as interfaces.

#### Example: `movies/service.go`

```go
package movies

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	CreateMovie(ctx context.Context, movie *Movie) error
	GetMovieByID(ctx context.Context, id uuid.UUID) (*Movie, error)
}

type service struct {
	repo Repository // Injected repository interface
}

// NewService accepts the Repository interface, maintaining loose coupling.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateMovie(ctx context.Context, movie *Movie) error {
	// Add business rules (validation, formatting) before calling repository
	return s.repo.CreateMovie(ctx, movie)
}

func (s *service) GetMovieByID(ctx context.Context, id uuid.UUID) (*Movie, error) {
	return s.repo.GetMovieByID(ctx, id)
}
```

#### Advanced Case: Service-to-Service Injection

If your `reservations` service needs to check if a `showtime` is valid, you inject the `showtimes.Service` interface into the `reservations.Service` constructor.

```go
package reservations

import (
	"context"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/showtimes"
)

type service struct {
	repo             Repository
	showtimeService  showtimes.Service // Injected service interface
}

func NewService(repo Repository, showtimeService showtimes.Service) Service {
	return &service{
		repo:            repo,
		showtimeService: showtimeService,
	}
}
```

---

### 3. Handler Layer (HTTP Request DI)

Handlers translate HTTP requests (Fiber context) into service calls. They accept the Service **interface**, alongside other cross-cutting concerns like loggers or validators.

We return concrete pointer types (e.g., `*Handler`) for handlers because they do not have alternative implementations (you don't mock a handler; you test it via HTTP request testing).

#### Example: `movies/handler.go`

```go
package movies

import (
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
)

type Handler struct {
	service Service         // Injected service interface
	logger  *zerolog.Logger // Injected logger pointer
}

// NewHandler acts as the constructor.
func NewHandler(service Service, logger *zerolog.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) CreateMovie(c fiber.Ctx) error {
	// Use h.logger.Info()... to log requests
	// Bind body to DTO and call h.service.CreateMovie(c.Context(), ...)
	return c.JSON(fiber.Map{"status": "created"})
}

func (h *Handler) GetMovieByID(c fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}
```

---

### 4. Router and Middleware Layer (Route DI)

Router functions register paths and link them to handlers. Middlewares also need dependency injection (for instance, an Auth middleware needs access to Redis for session verification, or a database connection to fetch users).

#### Custom Middleware DI with Closures

To inject dependencies into a middleware in Fiber, use a **closure factory** that accepts the dependencies and returns a `fiber.Handler`.

##### Example: `middleware/auth.go`

```go
package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

// Auth accepts its dependencies as arguments and returns a fiber.Handler closure.
func Auth(rdb *redis.Client) fiber.Handler {
	return func(c fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		// Perform session/token check using the injected Redis client
		val, err := rdb.Get(c.Context(), token).Result()
		if err != nil || val == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid session"})
		}

		return c.Next()
	}
}
```

##### Example: Registering Grouped Feature Routes (`movies/routes.go`)

```go
package movies

import (
	"github.com/gofiber/fiber/v3"
)

// RegisterRoutes accepts the Router group and the constructed Handler.
func RegisterRoutes(router fiber.Router, handler *Handler) {
	movies := router.Group("/movies")

	movies.Get("/:id", handler.GetMovieByID)
	movies.Post("/", handler.CreateMovie)
}
```

---

### 5. Dependency Orchestration Container

The `Dependencies` struct acts as your dependency container, compiling, instantiating, and wiring up the hierarchy top-to-bottom.

All infrastructure objects loaded in `main.go` are passed into the constructor function.

#### Example: `server/dependencies.go`

```go
package server

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/movies"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/config"
)

type Dependencies struct {
	MovieHandler *movies.Handler
	RedisClient  *redis.Client // We expose this if the middleware folder needs access to it
}

// BuildDependencies wires up every layer from Repository -> Service -> Handler.
func BuildDependencies(
	cfg *config.Config,
	pgPool *pgxpool.Pool,
	rdb *redis.Client,
	log *zerolog.Logger,
) *Dependencies {

	// 1. Repositories
	movieRepo := movies.NewRepository(pgPool, rdb)

	// 2. Services
	movieService := movies.NewService(movieRepo)

	// 3. Handlers
	movieHandler := movies.NewHandler(movieService, log)

	return &Dependencies{
		MovieHandler: movieHandler,
		RedisClient:  rdb,
	}
}
```

---

### 6. Linking to the Server & Global Router

We wire up our routes inside the global router configuration by passing the constructed dependency manager. This is also where middleware dependency injection takes place.

#### Example: `server/router.go`

```go
package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/movies"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/middleware"
)

func registerRoutes(app *fiber.App, deps *Dependencies) {
	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	v1 := app.Group("/api/v1")

	// Inject Redis into Auth middleware using the closure pattern
	protected := v1.Group("", middleware.Auth(deps.RedisClient))

	// Register sub-routes, injecting the MovieHandler dependency
	movies.RegisterRoutes(protected, deps.MovieHandler)
}
```

---

### 7. Entry Point Orchestration

Finally, the root of your application boots the database connections, initializes the configuration, starts the logger, and passes them down to generate the dependency chain.

#### Example: `cmd/api/main.go`

```go
func run() error {
	ctx := context.Background()

	// 1. Load config
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// 2. Initialize logger
	logInstance := logger.New(cfg.Env)

	// 3. Postgres
	pool, err := database.NewPostgresPool(ctx, cfg.PostgresDSN())
	if err != nil {
		return err
	}
	defer pool.Close()

	// 4. Redis
	rdb, err := database.NewRedisClient(ctx, cfg.RedisAddr())
	if err != nil {
		return err
	}
	defer rdb.Close()

	// 5. Build dependencies (DI orchestration step)
	deps := server.BuildDependencies(cfg, pool, rdb, &logInstance)

	// 6. Start server (Logger is passed into server configuration directly)
	srv := server.New(cfg, deps, &logInstance)
	return srv.Start()
}
```
