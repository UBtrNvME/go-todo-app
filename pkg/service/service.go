package service

import (
	"github.com/UBtrNvME/go-todo-app/pkg/domain"
	"github.com/UBtrNvME/go-todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoListInterface interface {
	Create(userID int, list domain.TodoList) (int, error)
	GetAll(userID int) ([]domain.TodoList, error)
	GetById(userID, listId int) (domain.TodoList, error)
	Update(userID, listId int, data interface{}) error
	Delete(userID, listId int) error
}

type TodoItemInterface interface {
}

type Service struct {
	Authorization
	TodoItemInterface
	TodoListInterface
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:     NewAuthService(repos.Authorization),
		TodoListInterface: NewTodoListService(repos.TodoListInterface),
	}
}
