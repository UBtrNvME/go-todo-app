package service

import (
	"github.com/UBtrNvME/go-todo-app/pkg/domain"
	"github.com/UBtrNvME/go-todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoListInterface
}

func NewTodoListService(repo repository.TodoListInterface) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s TodoListService) Create(userID int, list domain.TodoList) (int, error) {
	return s.repo.Create(userID, list)
}

func (s TodoListService) GetAll(userID int) ([]domain.TodoList, error) {
	return s.repo.GetAll(userID)
}
