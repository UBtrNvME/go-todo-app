package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {

}

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
	return &Repository{}
}
