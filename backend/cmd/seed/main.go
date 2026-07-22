package main

import (
	"context"
	"encoding/json"
	"errors" // added standard library errors package
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/config"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/database"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/migrations"
)

// TMDB API Response Structures
type TMDBGenre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TMDBGenreResponse struct {
	Genres []TMDBGenre `json:"genres"`
}

type TMDBMovie struct {
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	Overview     string  `json:"overview"`
	PosterPath   string  `json:"poster_path"`
	BackdropPath string  `json:"backdrop_path"`
	ReleaseDate  string  `json:"release_date"`
	VoteAverage  float64 `json:"vote_average"`
	GenreIDs     []int   `json:"genre_ids"`
	Runtime      int     `json:"runtime"` // fetched via detailed endpoint
}

type TMDBTrendingResponse struct {
	Results []TMDBMovie `json:"results"`
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("Seeding failed: %v", err)
	}
}

func run() error {
	ctx := context.Background()

	// 1. load configuration (including .env)
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	apiKey := os.Getenv("TMDB_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("TMDB_API_KEY environment variable is missing")
	}

	// 2. connect to PostgreSQL
	log.Println("Connecting to PostgreSQL...")
	pool, err := database.NewPostgresPool(ctx, cfg.PostgresDSN())
	if err != nil {
		return fmt.Errorf("postgres connection failed: %w", err)
	}
	defer pool.Close()

	// ensure database schema exists prior to seeding
	log.Println("Applying database migrations...")
	if err := migrations.RunMigrations(ctx, pool); err != nil {
		return fmt.Errorf("failed to apply migrations before seeding: %w", err)
	}

	// 3. fetch and Seed Genres
	log.Println("Fetching genres from TMDB...")
	tmdbGenres, err := fetchGenres(ctx, apiKey)
	if err != nil {
		return fmt.Errorf("failed to fetch TMDB genres: %w", err)
	}

	log.Println("Seeding genres into database...")
	genreMap, err := seedGenres(ctx, pool, tmdbGenres)
	if err != nil {
		return fmt.Errorf("failed to seed database genres: %w", err)
	}
	log.Printf("Successfully processed %d genres\n", len(genreMap))

	// 4. fetch Trending Movies List (Targeting 1000 items)
	log.Println("Fetching trending movies list from TMDB...")
	movies, err := fetchTrendingMovies(ctx, apiKey, 1000)
	if err != nil {
		return fmt.Errorf("failed to fetch trending movies: %w", err)
	}
	log.Printf("Retrieved %d trending movies from trending endpoints\n", len(movies))

	// 5. concurrently Fetch Movie Runtimes
	log.Println("Fetching detailed runtime information for movies...")
	enrichedMovies := fetchMovieDetailsConcurrently(ctx, apiKey, movies, 15) // 15 parallel workers

	// count inserted unique movies out of 1000
	uniqueIDs := make(map[int]bool)
	for _, m := range enrichedMovies {
		uniqueIDs[m.ID] = true
	}
	log.Printf("Parsed %d UNIQUE movies out of %d total items fetched\n", len(uniqueIDs), len(enrichedMovies))

	// 6. bulk Insert Movies & Associations inside a Transaction
	log.Println("Beginning transaction to insert movies...")
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	for i, m := range enrichedMovies {
		if err := insertMovie(ctx, tx, m, genreMap); err != nil {
			return fmt.Errorf("failed to insert movie %q (index %d): %w", m.Title, i, err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Println("Seeding completed successfully!")
	return nil
}

// fetchGenres fetches the full genre list mapping from TMDB.
func fetchGenres(ctx context.Context, apiKey string) ([]TMDBGenre, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/genre/movie/list?api_key=%s", apiKey)

	//nolint:gosec // g704: Trusted hardcoded destination base url
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	//nolint:gosec // g704: Trusting outbound client request to TMDB API domain
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var res TMDBGenreResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return res.Genres, nil
}

// seedGenres updates/inserts genres and maps TMDB IDs to DB UUIDs.
func seedGenres(ctx context.Context, db *pgxpool.Pool, genres []TMDBGenre) (map[int]uuid.UUID, error) {
	genreMap := make(map[int]uuid.UUID)

	for _, g := range genres {
		var existingID uuid.UUID
		// safe verification to fetch if already present or insert and return UUID
		err := db.QueryRow(ctx, "SELECT id FROM genres WHERE name = $1", g.Name).Scan(&existingID)
		if err == nil {
			genreMap[g.ID] = existingID
			continue
		}

		newID := uuid.New()
		_, err = db.Exec(ctx, "INSERT INTO genres (id, name) VALUES ($1, $2)", newID, g.Name)
		if err != nil {
			return nil, err
		}
		genreMap[g.ID] = newID
	}

	return genreMap, nil
}

// fetchTrendingMovies fetches target limits of trending movies from TMDB.
func fetchTrendingMovies(ctx context.Context, apiKey string, limit int) ([]TMDBMovie, error) {
	var movies []TMDBMovie
	client := &http.Client{Timeout: 8 * time.Second}

	// 20 items per page
	pagesNeeded := limit / 20
	if limit%20 != 0 {
		pagesNeeded++
	}

	for page := 1; page <= pagesNeeded; page++ {
		url := fmt.Sprintf("https://api.themoviedb.org/3/trending/movie/week?api_key=%s&page=%d", apiKey, page)

		//nolint:gosec // g704: Trusted hardcoded destination base url
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}

		//nolint:gosec // g704: Trusting outbound client request to TMDB API domain
		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("failed to reach TMDB API on page %d: %w", page, err)
		}

		var apiResp TMDBTrendingResponse
		err = json.NewDecoder(resp.Body).Decode(&apiResp)
		_ = resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("failed decoding page %d payload: %w", page, err)
		}

		movies = append(movies, apiResp.Results...)
		if len(apiResp.Results) == 0 {
			break
		}
	}

	if len(movies) > limit {
		movies = movies[:limit]
	}

	return movies, nil
}

// fetchMovieDetailsConcurrently fetches detailed information for each movie using worker routines.
func fetchMovieDetailsConcurrently(ctx context.Context, apiKey string, movies []TMDBMovie, workers int) []TMDBMovie {
	jobs := make(chan TMDBMovie, len(movies))
	results := make(chan TMDBMovie, len(movies))
	var wg sync.WaitGroup

	for w := 1; w <= workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := &http.Client{Timeout: 5 * time.Second}
			for m := range jobs {
				// handle potential context cancellation early
				select {
				case <-ctx.Done():
					m.Runtime = 120 // fallback runtime
					results <- m
					continue
				default:
				}

				url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%d?api_key=%s", m.ID, apiKey)

				//nolint:gosec // g704: Trusted hardcoded destination base url
				req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
				if err != nil {
					m.Runtime = 120 // fallback runtime
					results <- m
					continue
				}

				//nolint:gosec // g704: Trusting outbound client request to TMDB API domain
				resp, err := client.Do(req)
				if err != nil {
					m.Runtime = 120 // fallback runtime
					results <- m
					continue
				}

				var details struct {
					Runtime int `json:"runtime"`
				}
				if err := json.NewDecoder(resp.Body).Decode(&details); err == nil && details.Runtime > 0 {
					m.Runtime = details.Runtime
				} else {
					m.Runtime = 120 // fallback runtime
				}
				_ = resp.Body.Close()
				results <- m
			}
		}()
	}

	for _, m := range movies {
		jobs <- m
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	var enriched []TMDBMovie
	count := 0
	for em := range results {
		enriched = append(enriched, em)
		count++
		if count%100 == 0 {
			log.Printf("Enriched details for %d/%d movies...\n", count, len(movies))
		}
	}

	return enriched
}

// insertMovie inserts individual movies and configures their database genre associations.
func insertMovie(ctx context.Context, tx pgx.Tx, m TMDBMovie, genreMap map[int]uuid.UUID) error {
	var movieID uuid.UUID
	var releaseDate time.Time

	if m.ReleaseDate != "" {
		if parsed, err := time.Parse("2006-01-02", m.ReleaseDate); err == nil {
			releaseDate = parsed
		}
	}

	// 1. Check Movie existence inside the transaction safely
	err := tx.QueryRow(ctx, "SELECT id FROM movies WHERE tmdb_id = $1", m.ID).Scan(&movieID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// movie genuinely does not exist, safe to proceed with the insert inside the transaction
			movieID = uuid.New()
			insertQuery := `
				INSERT INTO movies (id, tmdb_id, title, overview, poster_path, backdrop_path, release_date, runtime, vote_average, created_at, updated_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW())`
			_, err = tx.Exec(ctx, insertQuery, movieID, m.ID, m.Title, m.Overview, m.PosterPath, m.BackdropPath, releaseDate, m.Runtime, m.VoteAverage)
			if err != nil {
				return fmt.Errorf("failed to execute insert statement: %w", err)
			}
		} else {
			// a real database error occurred (e.g. missing relation, bad column type). Return it directly.
			return fmt.Errorf("failed during existence lookup check: %w", err)
		}
	} else {
		// movie exists, perform the update safely
		updateQuery := `
			UPDATE movies SET 
				title = $1, overview = $2, poster_path = $3, backdrop_path = $4, 
				release_date = $5, runtime = $6, vote_average = $7, updated_at = NOW()
			WHERE id = $8`
		_, err = tx.Exec(ctx, updateQuery, m.Title, m.Overview, m.PosterPath, m.BackdropPath, releaseDate, m.Runtime, m.VoteAverage, movieID)
		if err != nil {
			return fmt.Errorf("failed to execute update statement: %w", err)
		}
	}

	// 2 clear old genre mappings to prevent obsolete constraints
	_, err = tx.Exec(ctx, "DELETE FROM movie_genres WHERE movie_id = $1", movieID)
	if err != nil {
		return fmt.Errorf("failed to clear obsolete movie genre mappings: %w", err)
	}

	// 3. map movie genres relationships
	for _, genreID := range m.GenreIDs {
		dbGenreID, exists := genreMap[genreID]
		if !exists {
			continue
		}
		_, err = tx.Exec(ctx, `
			INSERT INTO movie_genres (movie_id, genre_id) 
			VALUES ($1, $2) 
			ON CONFLICT DO NOTHING`, movieID, dbGenreID)
		if err != nil {
			return fmt.Errorf("failed to link movie to genre relation: %w", err)
		}
	}

	return nil
}
