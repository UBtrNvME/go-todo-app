package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/UBtrNvME/go-todo-app/pkg/domain"
	"github.com/UBtrNvME/go-todo-app/pkg/repository"
)
// Todo: move this into config or database
const salt = "kadfi83joalj03j0q0f09js"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) (string) {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}