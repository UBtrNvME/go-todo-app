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

func (s TodoListService) GetById(userID, listId int) (domain.TodoList, error) {
	return s.repo.GetById(userID, listId)
}

func (s TodoListService) Update(userID, listId int, data interface{}) error {
	panic("implement me")
}

func (s TodoListService) Delete(userID, listId int) error {
	panic("implement me")
}

