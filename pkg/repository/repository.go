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
		Authorization: NewAuthPostgres(db),
	}
}
