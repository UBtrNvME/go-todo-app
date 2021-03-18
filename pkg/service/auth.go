package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/UBtrNvME/go-todo-app/pkg/domain"
	"github.com/UBtrNvME/go-todo-app/pkg/repository"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"time"
)

// Todo: move this into config or database
const (
	salt       = "kadfi83joalj03j0q0f09js"
	tokenTTL   = 12 * time.Hour
	signingKey = "asdasdsadff849uji2hn0fj04GHFIUOBUIB982oilnLAHFOH98O"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	// GET users from DB
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	// Signing JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})
	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
