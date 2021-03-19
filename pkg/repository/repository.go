package repository

import (
	"github.com/UBtrNvME/go-todo-app/pkg/domain"
	"github.com/jmoiron/sqlx"
)

type (
	Authorization interface {
		CreateUser(user domain.User) (int, error)
		GetUser(username, password string) (domain.User, error)
	}
)

type TodoListInterface interface {
	Create(userID int, list domain.TodoList) (int, error)
	GetAll(userID int) ([]domain.TodoList, error)
	GetById(userID, listId int) (domain.TodoList, error)
	Update(userID, listId int, data interface{}) error
	Delete(userID, listId int) error
}

type TodoItemInterface interface {
}

type Repository struct {
	Authorization
	TodoItemInterface
	TodoListInterface
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:     NewAuthPostgres(db),
		TodoListInterface: NewTodoListPostgres(db),
	}
}
