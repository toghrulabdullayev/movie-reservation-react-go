package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/config"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/shared/hash"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/shared/jwtutil"
)

var _ Service = (*service)(nil)

type Service interface {
	Register(ctx context.Context, req RegisterRequest) (*User, error)
	Login(ctx context.Context, req LoginRequest) (*User, string, string, error)
	Refresh(ctx context.Context, refreshToken string) (*User, string, string, error)
	GetGoogleAuthURL() string
	HandleGoogleCallback(ctx context.Context, code string) (*User, string, string, error)
}

type service struct {
	repo Repository
	rdb  *redis.Client
	cfg  *config.Config
}

func NewService(repo Repository, rdb *redis.Client, cfg *config.Config) Service {
	return &service{
		repo: repo,
		rdb:  rdb,
		cfg:  cfg,
	}
}

func (s *service) Register(ctx context.Context, req RegisterRequest) (*User, error) {
	existingEmail, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if existingEmail != nil {
		return nil, errors.New("email is already taken")
	}

	existingUsername, err := s.repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if existingUsername != nil {
		return nil, errors.New("username is already taken")
	}

	hashed, err := hash.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		ID:           uuid.New(),
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hashed,
		Role:         "user",
	}

	if err := s.repo.CreateUser(ctx, u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *service) Login(ctx context.Context, req LoginRequest) (*User, string, string, error) {
	u, err := s.repo.GetUserByIdentity(ctx, req.Identity)
	if err != nil {
		return nil, "", "", err
	}
	if u == nil {
		return nil, "", "", errors.New("invalid credentials")
	}

	if u.PasswordHash == "" || !hash.CheckPasswordHash(req.Password, u.PasswordHash) {
		return nil, "", "", errors.New("invalid credentials")
	}

	accessToken, refreshToken, err := s.generateTokens(u)
	if err != nil {
		return nil, "", "", err
	}

	return u, accessToken, refreshToken, nil
}

func (s *service) Refresh(ctx context.Context, refreshToken string) (*User, string, string, error) {
	claims, err := jwtutil.ParseToken(refreshToken, s.cfg.JWTRefreshSecret)
	if err != nil {
		return nil, "", "", errors.New("invalid or expired refresh token")
	}

	u, err := s.repo.GetUserByID(ctx, claims.UserID)
	if err != nil {
		return nil, "", "", err
	}
	if u == nil {
		return nil, "", "", errors.New("user not found")
	}

	accessToken, newRefreshToken, err := s.generateTokens(u)
	if err != nil {
		return nil, "", "", err
	}

	return u, accessToken, newRefreshToken, nil
}

func (s *service) GetGoogleAuthURL() string {
	baseURL := "https://accounts.google.com/o/oauth2/v2/auth"
	vals := url.Values{}
	vals.Set("client_id", s.cfg.GoogleClientID)
	vals.Set("redirect_uri", s.cfg.GoogleRedirectURL)
	vals.Set("response_type", "code")
	vals.Set("scope", "https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email")
	vals.Set("access_type", "offline")
	vals.Set("prompt", "select_account")

	return fmt.Sprintf("%s?%s", baseURL, vals.Encode())
}

type googleTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type googleUserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (s *service) HandleGoogleCallback(ctx context.Context, code string) (*User, string, string, error) {
	// 1. Exchange Auth Code for Token
	tokenURL := "https://oauth2.googleapis.com/token"
	form := url.Values{}
	form.Set("code", code)
	form.Set("client_id", s.cfg.GoogleClientID)
	form.Set("client_secret", s.cfg.GoogleClientSecret)
	form.Set("redirect_uri", s.cfg.GoogleRedirectURL)
	form.Set("grant_type", "authorization_code")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, "", "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, "", "", fmt.Errorf("google token exchange returned status: %d", resp.StatusCode)
	}

	var tokenResp googleTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, "", "", err
	}

	// 2. Query User Info from Google Profile
	userInfoURL := "https://www.googleapis.com/oauth2/v2/userinfo"
	infoReq, err := http.NewRequestWithContext(ctx, http.MethodGet, userInfoURL, nil)
	if err != nil {
		return nil, "", "", err
	}
	infoReq.Header.Set("Authorization", "Bearer "+tokenResp.AccessToken)

	infoResp, err := http.DefaultClient.Do(infoReq)
	if err != nil {
		return nil, "", "", err
	}
	defer infoResp.Body.Close()

	if infoResp.StatusCode != http.StatusOK {
		return nil, "", "", fmt.Errorf("google profile fetch returned status: %d", infoResp.StatusCode)
	}

	var gUser googleUserInfo
	if err := json.NewDecoder(infoResp.Body).Decode(&gUser); err != nil {
		return nil, "", "", err
	}

	// 3. Resolve Database User matching Google Account Details
	u, err := s.repo.GetUserByGoogleID(ctx, gUser.ID)
	if err != nil {
		return nil, "", "", err
	}

	if u == nil {
		u, err = s.repo.GetUserByEmail(ctx, gUser.Email)
		if err != nil {
			return nil, "", "", err
		}

		if u != nil {
			u.GoogleID = &gUser.ID
			if err := s.repo.UpdateUser(ctx, u); err != nil {
				return nil, "", "", err
			}
		} else {
			username := gUser.Name
			if username == "" {
				username = strings.Split(gUser.Email, "@")[0]
			}
			existingU, err := s.repo.GetUserByUsername(ctx, username)
			if err != nil {
				return nil, "", "", err
			}
			if existingU != nil {
				username = fmt.Sprintf("%s_%s", username, s.randomStr(4))
			}

			randomBytes := make([]byte, 16)
			_, _ = rand.Read(randomBytes)
			unusableHash := "$2a$10$OauthPlaceholderUnusablePasswordHash" + hex.EncodeToString(randomBytes)

			u = &User{
				ID:           uuid.New(),
				Username:     username,
				Email:        gUser.Email,
				PasswordHash: unusableHash,
				Role:         "user",
				GoogleID:     &gUser.ID,
			}

			if err := s.repo.CreateUser(ctx, u); err != nil {
				return nil, "", "", err
			}
		}
	}

	accessToken, refreshToken, err := s.generateTokens(u)
	if err != nil {
		return nil, "", "", err
	}

	return u, accessToken, refreshToken, nil
}

func (s *service) generateTokens(u *User) (string, string, error) {
	accessToken, err := jwtutil.GenerateToken(u.ID, u.Role, s.cfg.JWTSecret, 15*time.Minute)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwtutil.GenerateToken(u.ID, u.Role, s.cfg.JWTRefreshSecret, 7*24*time.Hour)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *service) randomStr(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)[:n]
}
