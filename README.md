# Movie Reservation

A full-stack movie reservation system built with Go (Fiber) and React. The backend provides a RESTful API for managing movies, rooms, seats, showtimes, and reservations with JWT-based authentication and Google OAuth.

## Prerequisites

- Go 1.25+
- Docker & Docker Compose
- Make
- A [TMDB API key](https://www.themoviedb.org/documentation/api) (for seeding movie data)

## Getting Started

1. Start the infrastructure (PostgreSQL and Redis):

```sh
cd backend
docker compose up -d
```

2. Copy the example environment file and fill in your keys:

```sh
cp .env.example .env
```

3. Apply database migrations and seed movie data:

```sh
make migration-up
make seed
```

4. Start the development server:

```sh
make dev
```

The API will be available at `http://localhost:3000`.

## Swagger

Interactive API documentation is available at:

```
http://localhost:3000/swagger/
```

All endpoints are prefixed with `/api/v1`.

### Endpoints Overview

| Method             | Endpoint                       | Description                            |
| ------------------ | ------------------------------ | -------------------------------------- |
| **Auth**           |                                |                                        |
| `POST`             | `/api/v1/auth/register`        | Register a new user                    |
| `POST`             | `/api/v1/auth/login`           | Login with email/username and password |
| `POST`             | `/api/v1/auth/refresh`         | Refresh access token                   |
| `POST`             | `/api/v1/auth/logout`          | Logout and clear tokens                |
| `GET`              | `/api/v1/auth/me`              | Get current authenticated user         |
| `GET`              | `/api/v1/auth/google/login`    | Initiate Google OAuth login            |
| `GET`              | `/api/v1/auth/google/callback` | Google OAuth callback                  |
| **Movies**         |                                |                                        |
| `GET`              | `/api/v1/movies`               | List all movies (paginated)            |
| `GET`              | `/api/v1/movies/:id`           | Get movie by ID                        |
| **Rooms**          |                                |                                        |
| `GET`              | `/api/v1/rooms`                | List all rooms (paginated)             |
| `GET`              | `/api/v1/rooms/:id`            | Get room by ID                         |
| `POST`             | `/api/v1/rooms`                | Create a new room                      |
| `PUT`              | `/api/v1/rooms/:id`            | Update a room                          |
| `DELETE`           | `/api/v1/rooms/:id`            | Delete a room                          |
| **Seat Types**     |                                |                                        |
| `GET`              | `/api/v1/seats/types`          | List all seat types                    |
| `GET`              | `/api/v1/seats/types/:id`      | Get seat type by ID                    |
| `POST`             | `/api/v1/seats/types`          | Create a seat type                     |
| `PUT`              | `/api/v1/seats/types/:id`      | Update a seat type                     |
| `DELETE`           | `/api/v1/seats/types/:id`      | Delete a seat type                     |
| **Seats**          |                                |                                        |
| `GET`              | `/api/v1/seats`                | List all seats                         |
| `GET`              | `/api/v1/seats/:id`            | Get seat by ID                         |
| `POST`             | `/api/v1/seats`                | Create a new seat                      |
| `PUT`              | `/api/v1/seats/:id`            | Update a seat                          |
| `DELETE`           | `/api/v1/seats/:id`            | Delete a seat                          |
| **Showtimes**      |                                |                                        |
| `GET`              | `/api/v1/showtimes`            | List all showtimes                     |
| `GET`              | `/api/v1/showtimes/:id`        | Get showtime by ID                     |
| `POST`             | `/api/v1/showtimes`            | Create a showtime                      |
| `PUT`              | `/api/v1/showtimes/:id`        | Update a showtime                      |
| `DELETE`           | `/api/v1/showtimes/:id`        | Delete a showtime                      |
| **Showtime Seats** |                                |                                        |
| `GET`              | `/api/v1/showtimes/seats`      | List all showtime seats                |
| `GET`              | `/api/v1/showtimes/seats/:id`  | Get showtime seat by ID                |
| `POST`             | `/api/v1/showtimes/seats`      | Create a showtime seat                 |
| `PUT`              | `/api/v1/showtimes/seats/:id`  | Update a showtime seat                 |
| `DELETE`           | `/api/v1/showtimes/seats/:id`  | Delete a showtime seat                 |
| **Reservations**   |                                |                                        |
| `GET`              | `/api/v1/reservations`         | List all reservations (paginated)      |
| `GET`              | `/api/v1/reservations/:id`     | Get reservation by ID                  |
| `POST`             | `/api/v1/reservations`         | Create a reservation                   |
| `PUT`              | `/api/v1/reservations/:id`     | Update a reservation                   |
| `DELETE`           | `/api/v1/reservations/:id`     | Delete a reservation                   |

## Useful Commands

```sh
make help            # Show all available commands
make migration-up    # Apply database migrations
make migration-down  # Rollback last migration
make migration-new name=create_table  # Create a new migration
make seed            # Seed movies from TMDB
make swagger         # Regenerate Swagger docs
make dev             # Start dev server with hot reload
make lint            # Run linter
make test            # Run tests
make fmt             # Format code
```

## Tech Stack

- **Backend:** Go, Fiber v3, PostgreSQL, Redis, zerolog
- **Auth:** JWT + Google OAuth 2.0
- **API Docs:** Swagger (swaggo)
- **Migrations:** Goose
- **Seed Data:** TMDB API
